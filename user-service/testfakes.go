package main

import (
	"context"
	"errors"
	"net/http"
	"time"

	pb "github.com/willdot/go-do/user-service/proto/auth"
)

type fakeRepo struct {
	returnError bool
	users       []*pb.User
}

var errFake = errors.New("This is a fake error message")

func (f *fakeRepo) GetAll() ([]*pb.User, error) {

	if f.returnError {
		return nil, errFake
	}
	return f.users, nil
}

func (f *fakeRepo) Get(id string) (*pb.User, error) {

	if f.returnError {
		return nil, errFake
	}
	return &fakeUser, nil
}

func (f *fakeRepo) Create(user *pb.User) error {

	if f.returnError {
		return errFake
	}
	return nil
}

func (f *fakeRepo) GetByEmail(email string) (*pb.User, error) {

	if f.returnError {
		return nil, errFake
	}

	user := f.users[0]
	return user, nil
}

func (f *fakeRepo) Update(user *pb.User) error {

	if f.returnError {
		return errFake
	}
	return nil
}

func (f *fakeRepo) UpdatePassword(id, password string) error {

	if f.returnError {
		return errFake
	}
	return nil
}

var fakeUser = pb.User{
	Name:     "Fake",
	Email:    "fake@fake.com",
	Password: "$2a$10$cSOEkdxPPOrX8h/t3/Aw5e.vludnAzMGU38I3Cv0V/GAAwaqyJDaK",
	Company:  "fake",
}

func createService(returnError bool) userHandler {

	var users []*pb.User

	users = append(users, &fakeUser)

	fakeRepo := &fakeRepo{returnError, users}

	expireTime := time.Now().Add(time.Hour * 1).Unix()

	if returnError {
		expireTime = time.Now().AddDate(-1, 0, 0).Unix()
	}

	tokenService := TokenService{fakeRepo, expireTime}

	service := userHandler{fakeRepo, tokenService}

	return service
}

func createContext() context.Context {
	req, _ := http.NewRequest(http.MethodPost, "/", nil)

	ctx, _ := context.WithDeadline(req.Context(), time.Now())

	return ctx
}
