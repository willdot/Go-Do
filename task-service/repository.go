package main

import (
	taskPb "github.com/willdot/go-do/task-service/proto/task"
)

// Repository ..
type Repository interface {
	Get(userID string) ([]*taskPb.Task, error)
}
