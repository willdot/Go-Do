package main

import (
	"reflect"
	"testing"

	taskPb "github.com/willdot/go-do/task-service/proto/task"
)

func assertError(got, want error, t *testing.T) {
	if got != want {
		t.Errorf("got error '%v' but want error '%v'", got, want)
	}
}

func TestGetTasks(t *testing.T) {

	t.Run("get but returns an error in repo", func(t *testing.T) {
		service := createService(true, false)

		request := taskPb.Request{}
		response := taskPb.Response{}

		err := service.Get(createContext("t", true), &request, &response)

		assertError(err, errFake, t)
	})

	t.Run("get but returns an error in user service", func(t *testing.T) {
		service := createService(false, true)

		request := taskPb.Request{}
		response := taskPb.Response{}

		err := service.Get(createContext("t", true), &request, &response)

		assertError(err, errFake, t)
	})

	t.Run("get but returns an error for no token in metadata", func(t *testing.T) {
		service := createService(false, true)

		request := taskPb.Request{}
		response := taskPb.Response{}

		err := service.Get(createContext("", true), &request, &response)

		assertError(err, errFake, t)
	})

	t.Run("get but returns an error for metadata provided", func(t *testing.T) {
		service := createService(false, true)

		request := taskPb.Request{}
		response := taskPb.Response{}

		err := service.Get(createContext("", false), &request, &response)

		assertError(err, errNoMetaData, t)
	})

	t.Run("get all for user 1", func(t *testing.T) {

		service := createService(false, false)

		var want []*taskPb.Task

		want = append(want, &fakeTask1, &fakeTask2)

		request := taskPb.Request{}
		response := taskPb.Response{}

		err := service.Get(createContext("t", true), &request, &response)

		assertError(err, nil, t)

		if !reflect.DeepEqual(want, response.Tasks) {
			t.Errorf("want %v got %v", want, response.Tasks)
		}
	})

}
