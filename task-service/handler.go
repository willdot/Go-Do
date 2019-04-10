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

	suppliedToken, err := getTokenFromContext(ctx)

	if err != nil {
		return err
	}

	token := auth.Token{
		Token: suppliedToken,
	}

	validationResult, err := t.userClient.ValidateToken(ctx, &token)

	if err != nil {
		return err
	}

	userID := validationResult.UserId

	tasks, err := t.repo.Get(userID)

	if err != nil {
		return err
	}

	res.Tasks = tasks

	return nil
}

func getTokenFromContext(ctx context.Context) (string, error) {

	meta, ok := metadata.FromContext(ctx)

	if !ok {
		return "", errNoMetaData
	}

	token := meta["Token"]

	return token, nil
}
