package main

import (
	"errors"
	"fmt"
	"log"

	pb "github.com/willdot/go-do/user-service/proto/auth"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

var errUnknownUser = errors.New("User not found")

type userHandler struct {
	repo         Repository
	tokenService TokenService
}

func (u *userHandler) Create(ctx context.Context, req *pb.User, res *pb.Response) error {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}

	req.Password = string(hashedPass)

	err = u.repo.Create(req)

	if err != nil {
		return err
	}

	res.User = req

	return nil
}

func (u *userHandler) Get(ctx context.Context, req *pb.User, res *pb.Response) error {

	user, err := u.repo.Get(req.Id)

	if err != nil {
		return err
	}

	res.User = user
	return nil
}

func (u *userHandler) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {

	users, err := u.repo.GetAll()

	if err != nil {
		return err
	}

	res.Users = users

	return nil
}

func (u *userHandler) Update(ctx context.Context, req *pb.User, res *pb.Response) error {
	err := u.repo.Update(req)

	if err != nil {
		return err
	}

	res.User = req

	return nil
}

func (u *userHandler) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {

	log.Println("Logging in with: ", req.Email, req.Password)

	user, err := u.repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := u.tokenService.Encode(user)

	if err != nil {
		return err
	}

	res.Token = token
	return nil
}

func (u *userHandler) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {

	claims, err := u.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errUnknownUser
	}

	res.Valid = true

	return nil
}

func (u *userHandler) ChangePassword(ctx context.Context, req *pb.PasswordChange, res *pb.Token) error {

	return nil
}
