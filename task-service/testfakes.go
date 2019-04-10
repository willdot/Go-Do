package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/micro/go-micro/client"
	pb "github.com/willdot/go-do/task-service/proto/task"
	auth "github.com/willdot/go-do/user-service/proto/auth"
	"golang.org/x/net/context"
)

var errFake = errors.New("This is a fake error message")

type fakeRepo struct {
	returnError bool
	tasks       []*pb.Task
}

func (f *fakeRepo) Get(userId string) ([]*pb.Task, error) {

	if f.returnError {
		return nil, errFake
	}

	return f.tasks, nil
}

var fakeTask = pb.Task{
	Id:          "123",
	Title:       "Test",
	Description: "Do something",
	UserId:      "111",
	CreatedDate: 1,
	DailyDo:     false,
}

func createService(returnError bool) taskHandler {

	var tasks []*pb.Task

	tasks = append(tasks, &fakeTask)

	fakeRepo := &fakeRepo{returnError, tasks}

	fakeAuthClient := &fakeUserHandler{}

	service := taskHandler{fakeRepo, fakeAuthClient}

	return service
}

func createContext(token string) context.Context {
	req, _ := http.NewRequest(http.MethodPost, "/", nil)

	ctx, _ := context.WithDeadline(req.Context(), time.Now())

	return ctx
}

// CustomClaims ..
type CustomClaims struct {
	User *auth.User
}

// Authable ..
type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *auth.User) (string, error)
}

type fakeUserHandler struct {
	fakeTokenService Authable
}

func (u *fakeUserHandler) Create(ctx context.Context, req *auth.User, opts ...client.CallOption) (*auth.Response, error) {
	return nil, nil
}

func (u *fakeUserHandler) Get(ctx context.Context, req *auth.User, opts ...client.CallOption) (*auth.Response, error) {
	return nil, nil
}

func (u *fakeUserHandler) GetAll(ctx context.Context, req *auth.Request, opts ...client.CallOption) (*auth.Response, error) {
	return nil, nil

}

func (u *fakeUserHandler) Update(ctx context.Context, req *auth.User, opts ...client.CallOption) (*auth.Response, error) {
	return nil, nil
}

func (u *fakeUserHandler) Auth(ctx context.Context, req *auth.User, opts ...client.CallOption) (*auth.Token, error) {
	return nil, nil

}

func (u *fakeUserHandler) ValidateToken(ctx context.Context, req *auth.Token, opts ...client.CallOption) (*auth.Token, error) {

	token := &auth.Token{
		UserId: "123",
	}

	return token, nil
}

func (u *fakeUserHandler) ChangePassword(ctx context.Context, req *auth.PasswordChange, opts ...client.CallOption) (*auth.Token, error) {
	return nil, nil
}
