package main

import (
	"reflect"
	"testing"

	pb "github.com/willdot/go-do/task-service/proto/task"
)

func TestGetTasks(t *testing.T) {

	t.Run("get all and returns all tasks", func(t *testing.T) {
		service := createService(false)

		var want []*pb.Task

		want = append(want, &fakeTask)

		request := pb.Request{}
		response := pb.Response{}

		err := service.Get(createContext(""), &request, &response)

		if err != nil {
			t.Errorf("wanted %v but got %v", nil, err)
		}

		if !reflect.DeepEqual(want, response.Tasks) {
			t.Errorf("want %v got %v", want, response.Tasks)
		}
	})

	t.Run("get all for user", func(t *testing.T) {

		service := createService(false)

		var want []*pb.Task

		want = append(want, &fakeTask)

		request := pb.Request{}
		response := pb.Response{}

		err := service.Get(createContext(""), &request, &response)

		if err != nil {
			t.Errorf("wanted %v but got %v", nil, err)
		}

		if !reflect.DeepEqual(want, response.Tasks) {
			t.Errorf("want %v got %v", want, response.Tasks)
		}
	})

	t.Run("get but returns an error", func(t *testing.T) {
		service := createService(true)

		request := pb.Request{}
		response := pb.Response{}

		err := service.Get(createContext(""), &request, &response)

		if err != errFake {
			t.Errorf("wanted %v but got %v", nil, err)
		}
	})
}
