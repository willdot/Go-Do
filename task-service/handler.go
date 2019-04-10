package main

import (
	"golang.org/x/net/context"

	pb "github.com/willdot/go-do/task-service/proto/task"
	auth "github.com/willdot/go-do/user-service/proto/auth"
)

type taskHandler struct {
	repo       Repository
	userClient auth.AuthClient
}

func (t *taskHandler) Get(ctx context.Context, req *pb.Request, res *pb.Response) error {

	token := auth.Token{
		Token: "",
	}
	validationResult, err := t.userClient.ValidateToken(ctx, &token)

	userID := validationResult.UserId

	tasks, err := t.repo.Get(userID)

	if err != nil {
		return err
	}

	res.Tasks = tasks

	return nil
}
