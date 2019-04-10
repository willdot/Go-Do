package main

import (
	"errors"

	"golang.org/x/net/context"

	"github.com/micro/go-micro/metadata"
	pb "github.com/willdot/go-do/task-service/proto/task"
	auth "github.com/willdot/go-do/user-service/proto/auth"
)

var errNoMetaData = errors.New("no auth meta data found in request")

type taskHandler struct {
	repo       Repository
	userClient auth.AuthClient
}

func (t *taskHandler) Get(ctx context.Context, req *pb.Request, res *pb.Response) error {

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

// so that we can get the user id to get tasks for, we get the supplied token, validate it,
// and then get the user id. This means not having to send the user id in the request, which
// limits the chance of random api calls being made with guessed user id
func (t *taskHandler) getUserIDFromTokenInContext(ctx context.Context) (string, error) {
	suppliedToken, err := getTokenFromContext(ctx)

	if err != nil {
		return "", err
	}

	token := auth.Token{
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
