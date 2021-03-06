package main

import (
	"errors"
	"fmt"
	"log"

	authPb "github.com/willdot/go-do/user-service/proto/auth"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

var errUnknownUser = errors.New("User not found")

var errTokenPasswordNotValid = errors.New("Token password no longer valid")

type userHandler struct {
	repo         Repository
	tokenService TokenService
}

func (u *userHandler) Create(ctx context.Context, req *authPb.User, res *authPb.Response) error {

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

func (u *userHandler) Get(ctx context.Context, req *authPb.User, res *authPb.Response) error {

	user, err := u.repo.Get(req.Id)

	if err != nil {
		return err
	}

	res.User = user
	return nil
}

func (u *userHandler) GetAll(ctx context.Context, req *authPb.Request, res *authPb.Response) error {

	users, err := u.repo.GetAll()

	if err != nil {
		return err
	}

	res.Users = users

	return nil
}

func (u *userHandler) Update(ctx context.Context, req *authPb.User, res *authPb.Response) error {
	err := u.repo.Update(req)

	if err != nil {
		return err
	}

	res.User = req

	return nil
}

func (u *userHandler) Auth(ctx context.Context, req *authPb.User, res *authPb.Token) error {

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

func (u *userHandler) ValidateToken(ctx context.Context, req *authPb.Token, res *authPb.Token) error {

	claims, err := u.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errUnknownUser
	}

	// Get the user from the repo so that we can compare the password in the token with the users actual password.
	// Otherwise they may be using a token that was assigned before a password change
	user, err := u.repo.Get(claims.User.Id)

	if err != nil {
		return err
	}

	if user.Password != claims.User.Password {
		return errTokenPasswordNotValid
	}

	res.Valid = true
	res.UserId = claims.User.Id

	return nil
}

func (u *userHandler) ChangePassword(ctx context.Context, req *authPb.PasswordChange, res *authPb.Token) error {

	user, err := u.repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return err
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}

	user.Password = string(hashedPass)

	err = u.repo.UpdatePassword(user.Id, user.Password)

	if err != nil {
		return err
	}

	token, err := u.tokenService.Encode(user)

	if err != nil {
		return err
	}

	res.Token = token
	return nil
}
