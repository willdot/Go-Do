package main

import (
	"net/http"
	"reflect"
	"testing"
	"time"

	"golang.org/x/net/context"

	pb "github.com/willdot/go-do/task-service/proto/task"
)

func TestGetTasks(t *testing.T) {

	service := createService(false)

	var want []*pb.Task

	want = append(want, &fakeTask)

	request := pb.Request{}
	response := pb.Response{}

	err := service.Get(createContext(), &request, &response)

	if err != nil {
		t.Errorf("wanted %v but got %v", nil, err)
	}

	if !reflect.DeepEqual(want, response.Tasks) {
		t.Errorf("want %v got %v", want, response.Tasks)
	}

}

type fakeRepo struct {
	returnError bool
	tasks       []*pb.Task
}

func (f *fakeRepo) Get() ([]*pb.Task, error) {
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

	service := taskHandler{fakeRepo}

	return service
}

func createContext() context.Context {
	req, _ := http.NewRequest(http.MethodPost, "/", nil)

	ctx, _ := context.WithDeadline(req.Context(), time.Now())

	return ctx
}
