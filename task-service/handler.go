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
var errDailyDoAlreadyExists = errors.New("There is already a task set as daily do")

type taskHandler struct {
	repo       Repository
	userClient authPb.AuthClient
}

// Get satisfies the Get RPC for the Task proto and gets tasks for a user
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

// Create satisfies the Create RPC for the Task proto and creates a new task for a user
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

// Update satisfies the Update RPC for the Task proto and updates a task for a user
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

// ChangeDailyDoStatus satisfies the ChangeDailyDoStatus RPC for the Task proto and sets the daily do status of a task
func (t *taskHandler) ChangeDailyDoStatus(ctx context.Context, req *taskPb.DailyDoStatusRequest, res *taskPb.Response) error {

	userID, err := t.getUserIDFromTokenInContext(ctx)

	if err != nil {
		return err
	}

	// If setting a daily do, check there isn't already a daily do set
	if req.Status {
		existingDailyDo, err := t.repo.GetDailyDoForUser(userID)

		if err != nil {
			return err
		}

		if existingDailyDo != nil {
			return errDailyDoAlreadyExists
		}
	}

	task := taskPb.Task{
		Id:      req.TaskId,
		UserId:  userID,
		DailyDo: req.Status,
	}

	err = t.repo.SetDailyDoStatus(&task)

	if err != nil {
		return err
	}

	return nil
}

// CompleteTask sets the task CompletedDate
func (t *taskHandler) CompleteTask(ctx context.Context, req *taskPb.CompleteTaskRequest, res *taskPb.Response) error {

	userID, err := t.getUserIDFromTokenInContext(ctx)

	if err != nil {
		return err
	}

	task := taskPb.Task{
		Id:     req.TaskId,
		UserId: userID,
	}

	if req.Completed {
		task.CompletedDate = int64(time.Now().Unix())
		task.DailyDo = false
	} else {
		task.CompletedDate = 0
	}

	err = t.repo.CompleteTask(&task)

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
