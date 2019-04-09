package main

import (
	pb "go-do/user-service/proto/auth"

	"golang.org/x/net/context"
)

type userHandler struct {
}

func (u *userHandler) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	return nil
}

func (u *userHandler) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	return nil
}

func (u *userHandler) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	return nil
}

func (u *userHandler) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	return nil
}

func (u *userHandler) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}

func (u *userHandler) Update(ctx context.Context, req *pb.User, res *pb.Response) error {
	return nil
}
