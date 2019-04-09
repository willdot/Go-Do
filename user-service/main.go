package main

import (
	pb "go-do/user-service/proto/auth"

	"github.com/micro/go-micro"
)

func main() {

	srv := micro.NewService(
		micro.Name("go_do.auth"),
	)

	srv.Init()

	pb.RegisterAuthHandler(srv.Server(), &userHandler{})
}
