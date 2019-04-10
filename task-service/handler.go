package main

import (
	"golang.org/x/net/context"

	pb "github.com/willdot/go-do/task-service/proto/task"
)

type taskHandler struct {
	repo Repository
}

func (t *taskHandler) Get(ctx context.Context, req *pb.Request, res *pb.Response) error {

	tasks, err := t.repo.Get()

	if err != nil {
		return err
	}

	res.Tasks = tasks

	return nil
}
