package main

import (
	"fmt"
	"time"

	pb "github.com/willdot/go-do/user-service/proto/auth"

	"github.com/micro/go-micro"
)

func main() {

	CassandraSession := Session
	defer CassandraSession.Close()

	tokenExpireTime := time.Now().Add(time.Hour * 72).Unix()

	repo := &UserRepository{CassandraSession}

	tokenService := TokenService{repo, tokenExpireTime}

	srv := micro.NewService(
		micro.Name("go_do.auth"),
	)

	srv.Init()

	pb.RegisterAuthHandler(srv.Server(), &userHandler{repo, tokenService})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
