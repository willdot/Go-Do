package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	taskPb "github.com/willdot/go-do/task-service/proto/task"
	authPb "github.com/willdot/go-do/user-service/proto/auth"
	"golang.org/x/net/context"
)

const (
	userID1 = "111"
	userID2 = "222"
)

var errFake = errors.New("This is a fake error message")

type fakeRepo struct {
	// returnError is used as a flag to return a fake error
	returnError bool
	tasks       []*taskPb.Task
}

func (f *fakeRepo) Get(userID string) ([]*taskPb.Task, error) {

	if f.returnError {
		return nil, errFake
	}

	var tasks []*taskPb.Task

	for _, v := range f.tasks {
		if v.UserId == userID {
			tasks = append(tasks, v)
		}
	}

	return tasks, nil
}

func (f *fakeRepo) Create(task *taskPb.Task) error {
	if f.returnError {
		return errFake
	}

	task.Id = "123"

	return nil
}

func (f *fakeRepo) Update(task *taskPb.Task) error {
	if f.returnError {
		return errFake
	}

	var taskToUpdate *taskPb.Task

	for _, v := range f.tasks {
		if v.Id == task.Id {
			taskToUpdate = v
			break
		}
	}

	if taskToUpdate == nil {
		return errTaskNotFound
	}

	if taskToUpdate.UserId != task.UserId {
		return errTaskUserIDNotMatched
	}

	taskToUpdate.Title = task.Title
	taskToUpdate.Description = task.Description

	return nil
}

var fakeTask1 = taskPb.Task{
	Id:          "123",
	Title:       "Test1",
	Description: "Do something",
	UserId:      userID1,
	CreatedDate: 1,
	DailyDo:     false,
}

var fakeTask2 = taskPb.Task{
	Id:          "123",
	Title:       "Test2",
	Description: "Do something",
	UserId:      userID1,
	CreatedDate: 1,
	DailyDo:     false,
}

var fakeTask3 = taskPb.Task{
	Id:          "123",
	Title:       "Test3",
	Description: "Do something",
	UserId:      userID2,
	CreatedDate: 1,
	DailyDo:     false,
}

// createService creates a fake service with mocks.
func createService(repoReturnError, userHandlerReturnError, userIDInTokenMatchesTask bool) taskHandler {

	var tasks []*taskPb.Task

	tasks = append(tasks, &fakeTask1, &fakeTask2, &fakeTask3)

	fakeRepo := &fakeRepo{repoReturnError, tasks}

	fakeAuthClient := &fakeUserHandler{userHandlerReturnError, userIDInTokenMatchesTask}

	service := taskHandler{fakeRepo, fakeAuthClient}

	return service
}

// createContext creates a fake context and adds in a token if required
func createContext(token string, addMetaData bool) context.Context {
	req, _ := http.NewRequest(http.MethodPost, "/", nil)

	ctx, cancel := context.WithDeadline(req.Context(), time.Now())

	if addMetaData {
		tokenMap := make(map[string]string)

		if token != "" {
			tokenMap["token"] = token
		}
		ctx = metadata.NewContext(ctx, tokenMap)
	}

	// error given when cancel func returned from above call is discarded, so just defering it to remove the error. Has no impact on a test
	defer cancel()

	return ctx
}

// This is used so that the validate token method can used to get the user id from a JWT
type fakeUserHandler struct {
	// ReturnError is a flag to simulate an error in the method
	returnError bool
	// UseIdMatches is a flag to set the user id found inside the JWT. Then in the repo when a comparison between the user id
	// of the task that has been sent in the request doesn't match the user id in the JWT, an error can be returned
	userIDMatches bool
}

func (u *fakeUserHandler) Create(ctx context.Context, req *authPb.User, opts ...client.CallOption) (*authPb.Response, error) {
	return nil, nil
}

func (u *fakeUserHandler) Get(ctx context.Context, req *authPb.User, opts ...client.CallOption) (*authPb.Response, error) {
	return nil, nil
}

func (u *fakeUserHandler) GetAll(ctx context.Context, req *authPb.Request, opts ...client.CallOption) (*authPb.Response, error) {
	return nil, nil

}

func (u *fakeUserHandler) Update(ctx context.Context, req *authPb.User, opts ...client.CallOption) (*authPb.Response, error) {
	return nil, nil
}

func (u *fakeUserHandler) Auth(ctx context.Context, req *authPb.User, opts ...client.CallOption) (*authPb.Token, error) {
	return nil, nil

}

func (u *fakeUserHandler) ValidateToken(ctx context.Context, req *authPb.Token, opts ...client.CallOption) (*authPb.Token, error) {

	if u.returnError {
		return nil, errFake
	}

	userID := userID1

	if !u.userIDMatches {
		userID = userID2
	}

	token := &authPb.Token{
		UserId: userID,
	}

	return token, nil
}

func (u *fakeUserHandler) ChangePassword(ctx context.Context, req *authPb.PasswordChange, opts ...client.CallOption) (*authPb.Token, error) {
	return nil, nil
}
