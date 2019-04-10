package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	pb "github.com/willdot/go-do/task-service/proto/task"
	auth "github.com/willdot/go-do/user-service/proto/auth"
	"golang.org/x/net/context"
)

var errFake = errors.New("This is a fake error message")

type fakeRepo struct {
	// returnError is used as a flag to return a fake error
	returnError bool
	tasks       []*pb.Task
}

func (f *fakeRepo) Get(userID string) ([]*pb.Task, error) {

	if f.returnError {
		return nil, errFake
	}

	var tasks []*pb.Task

	for _, v := range f.tasks {
		if v.UserId == userID {
			tasks = append(tasks, v)
		}
	}

	return tasks, nil
}

var fakeTask1 = pb.Task{
	Id:          "123",
	Title:       "Test1",
	Description: "Do something",
	UserId:      "111",
	CreatedDate: 1,
	DailyDo:     false,
}

var fakeTask2 = pb.Task{
	Id:          "123",
	Title:       "Test2",
	Description: "Do something",
	UserId:      "111",
	CreatedDate: 1,
	DailyDo:     false,
}

var fakeTask3 = pb.Task{
	Id:          "123",
	Title:       "Test3",
	Description: "Do something",
	UserId:      "222",
	CreatedDate: 1,
	DailyDo:     false,
}

func createService(repoReturnError, userHandlerReturnError bool) taskHandler {

	var tasks []*pb.Task

	tasks = append(tasks, &fakeTask1, &fakeTask2, &fakeTask3)

	fakeRepo := &fakeRepo{repoReturnError, tasks}

	fakeAuthClient := &fakeUserHandler{userHandlerReturnError}

	service := taskHandler{fakeRepo, fakeAuthClient}

	return service
}

func createContext(token string, addMetaData bool) context.Context {
	req, _ := http.NewRequest(http.MethodPost, "/", nil)

	ctx, _ := context.WithDeadline(req.Context(), time.Now())

	if addMetaData {
		tokenMap := make(map[string]string)

		if token != "" {
			tokenMap["token"] = token
		}
		ctx = metadata.NewContext(ctx, tokenMap)
	}

	return ctx
}

type fakeUserHandler struct {
	returnError bool
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

	if u.returnError {
		return nil, errFake
	}
	token := &auth.Token{
		UserId: "111",
	}

	return token, nil
}

func (u *fakeUserHandler) ChangePassword(ctx context.Context, req *auth.PasswordChange, opts ...client.CallOption) (*auth.Token, error) {
	return nil, nil
}
