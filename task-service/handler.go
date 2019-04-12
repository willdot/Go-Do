package main

import (
	"errors"
	"time"

	"golang.org/x/net/context"

	"github.com/micro/go-micro/metadata"
	taskPb "github.com/willdot/go-do/task-service/proto/task"
	authPb "github.com/willdot/go-do/user-service/proto/auth"
)

var errNoMetaData = errors.New("no auth meta data found in request")

type taskHandler struct {
	repo       Repository
	userClient authPb.AuthClient
}

func (t *taskHandler) Get(ctx context.Context, req *taskPb.Request, res *taskPb.Response) error {

	userID, err := t.getUserIDFromTokenInContext(ctx)

	if err != nil {
		return err
	}
	tasks, err := t.repo.Get(userID)

	if err != nil {
		return err
	}

	res.Tasks = tasks

	return nil
}

func (t *taskHandler) Create(ctx context.Context, req *taskPb.CreateTask, res *taskPb.Response) error {

	userID, err := t.getUserIDFromTokenInContext(ctx)

	if err != nil {
		return err
	}

	task := taskPb.Task{
		Title:       req.Title,
		Description: req.Description,
		DailyDo:     req.DailyDo,
		UserId:      userID,
		CreatedDate: int64(time.Now().Unix()),
	}

	err = t.repo.Create(&task)

	if err != nil {
		return err
	}

	res.Task = &task

	return nil
}

func (t *taskHandler) Update(ctx context.Context, req *taskPb.UpdateTask, res *taskPb.Response) error {

	userID, err := t.getUserIDFromTokenInContext(ctx)

	if err != nil {
		return err
	}

	task := taskPb.Task{
		Id:          req.TaskId,
		Title:       req.Title,
		Description: req.Description,
		UserId:      userID,
	}

	err = t.repo.Update(&task)

	if err != nil {
		return err
	}

	return nil
}

// so that we can get the user id to use on the functions, we get the supplied token, validate it,
// and then get the user id. This means not having to send the user id in the request, which
// limits the chance of random api calls being made with guessed user id
func (t *taskHandler) getUserIDFromTokenInContext(ctx context.Context) (string, error) {
	suppliedToken, err := getTokenFromContext(ctx)

	if err != nil {
		return "", err
	}

	token := authPb.Token{
		Token: suppliedToken,
	}

	validationResult, err := t.userClient.ValidateToken(ctx, &token)

	if err != nil {
		return "", err
	}

	return validationResult.UserId, nil
}

func getTokenFromContext(ctx context.Context) (string, error) {

	meta, ok := metadata.FromContext(ctx)

	if !ok {
		return "", errNoMetaData
	}

	token := meta["Token"]

	return token, nil
}
