package main

import (
	pb "github.com/willdot/go-do/task-service/proto/task"
)

// Repository ..
type Repository interface {
	Get() ([]*pb.Task, error)
}
