package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	taskPb "github.com/willdot/go-do/task-service/proto/task"
	authPb "github.com/willdot/go-do/user-service/proto/auth"
	"golang.org/x/net/context"
)

func main() {

	CassandraSession := Session
	defer CassandraSession.Close()

	repo := &TaskRepository{CassandraSession}

	srv := micro.NewService(
		micro.Name("go_do.task"),
		micro.WrapHandler(AuthWrapper),
	)

	authClient := authPb.NewAuthClient("go_do.auth", srv.Client())

	srv.Init()

	taskPb.RegisterTaskServiceHandler(srv.Server(), &taskHandler{repo, authClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

// AuthWrapper is a wrapper for authorising using a JWT
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)

		if !ok {
			return errors.New("no auth meta data found in request")
		}

		token := meta["Token"]
		log.Println("Authenticating token: ", token)

		authClient := authPb.NewAuthClient("go_do.auth", client.DefaultClient)

		authResp, err := authClient.ValidateToken(ctx, &authPb.Token{
			Token: token,
		})

		log.Println("Auth response: ", authResp)
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}
