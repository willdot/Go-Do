package main

import (
	"fmt"

	"github.com/micro/go-micro"
	taskPb "github.com/willdot/go-do/task-service/proto/task"
)

func main() {

	srv := micro.NewService(
		micro.Name("go_do.task"),
	)

	srv.Init()

	taskPb.RegisterTaskServiceHandler(srv.Server(), nil)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
