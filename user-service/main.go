package main

import (
	"fmt"
	pb "github.com/willdot/go-do/user-service/proto/auth"

	"github.com/micro/go-micro"
)

func main() {

	srv := micro.NewService(
		micro.Name("go_do.auth"),
	)

	srv.Init()

	pb.RegisterAuthHandler(srv.Server(), &userHandler{})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
