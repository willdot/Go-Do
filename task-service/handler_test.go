package main

import (
	"reflect"
	"testing"
	"time"

	taskPb "github.com/willdot/go-do/task-service/proto/task"
)

func assertError(got, want error, t *testing.T) {
	if got != want {
		t.Errorf("got error '%v' but want error '%v'", got, want)
	}
}

func TestGetTasks(t *testing.T) {

	t.Run("get but returns an error in repo", func(t *testing.T) {
		service := createService(true, false, true)

		request := taskPb.Request{}
		response := taskPb.Response{}

		err := service.Get(createContext("t", true), &request, &response)

		assertError(err, errFake, t)
	})

	t.Run("get but returns an error for no token in metadata", func(t *testing.T) {
		service := createService(false, true, true)

		request := taskPb.Request{}
		response := taskPb.Response{}

		err := service.Get(createContext("", true), &request, &response)

		assertError(err, errFake, t)
	})

	t.Run("get but returns an error for metadata provided", func(t *testing.T) {
		service := createService(false, true, true)

		request := taskPb.Request{}
		response := taskPb.Response{}

		err := service.Get(createContext("", false), &request, &response)

		assertError(err, errNoMetaData, t)
	})

	t.Run("get all for user 1", func(t *testing.T) {

		service := createService(false, false, true)

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

func TestCreateTasks(t *testing.T) {

	t.Run("create but returns an error in repo", func(t *testing.T) {
		service := createService(true, false, true)

		request := taskPb.CreateTask{}
		response := taskPb.Response{}

		err := service.Create(createContext("t", true), &request, &response)

		assertError(err, errFake, t)
	})

	t.Run("create but returns an error for no token in metadata", func(t *testing.T) {
		service := createService(false, true, true)

		request := taskPb.CreateTask{}
		response := taskPb.Response{}

		err := service.Create(createContext("", true), &request, &response)

		assertError(err, errFake, t)
	})

	t.Run("create but returns an error for metadata provided", func(t *testing.T) {
		service := createService(false, true, true)

		request := taskPb.CreateTask{}
		response := taskPb.Response{}

		err := service.Create(createContext("", false), &request, &response)

		assertError(err, errNoMetaData, t)
	})

	t.Run("create task for user 1", func(t *testing.T) {

		service := createService(false, false, true)

		want := "123"

		request := taskPb.CreateTask{
			Title:       "fake",
			Description: "fake",
			DailyDo:     false,
		}

		response := taskPb.Response{}

		err := service.Create(createContext("t", true), &request, &response)

		assertError(err, nil, t)

		if want != response.Task.Id {
			t.Errorf("want %v got %v", want, response.Task.Id)
		}

		currentDate := time.Now().Format("Jan 2 15:04")
		createdDate := time.Unix(response.Task.CreatedDate, 0).Format("Jan 2 15:04")

		if currentDate != createdDate {
			t.Errorf("want %v got %v", currentDate, createdDate)
		}

	})
}

func TestUpdateTask(t *testing.T) {

	t.Run("update but repo returns error", func(t *testing.T) {

		service := createService(true, false, true)

		request := taskPb.UpdateTask{}
		response := taskPb.Response{}

		err := service.Update(createContext("t", true), &request, &response)

		assertError(err, errFake, t)

	})

	t.Run("update but returns an error for no token in metadata", func(t *testing.T) {
		service := createService(false, true, true)

		request := taskPb.UpdateTask{}
		response := taskPb.Response{}

		err := service.Update(createContext("", true), &request, &response)

		assertError(err, errFake, t)
	})

	t.Run("update but returns an error for metadata provided", func(t *testing.T) {
		service := createService(false, true, true)

		request := taskPb.UpdateTask{}
		response := taskPb.Response{}

		err := service.Update(createContext("", false), &request, &response)

		assertError(err, errNoMetaData, t)
	})

	t.Run("task not found", func(t *testing.T) {
		service := createService(false, false, true)

		request := taskPb.UpdateTask{
			TaskId: "not found",
		}
		response := taskPb.Response{}

		err := service.Update(createContext("t", true), &request, &response)

		assertError(err, errTaskNotFound, t)
	})

	t.Run("tasks user id doesn't match id in token", func(t *testing.T) {
		service := createService(false, false, false)

		request := taskPb.UpdateTask{
			TaskId: "123",
		}
		response := taskPb.Response{}

		err := service.Update(createContext("t", true), &request, &response)

		assertError(err, errTaskUserIDNotMatched, t)
	})

	t.Run("update fakeTask1 for user 1", func(t *testing.T) {

		service := createService(false, false, true)

		request := taskPb.UpdateTask{
			TaskId:      "123",
			Title:       "new",
			Description: "new",
		}

		response := taskPb.Response{}

		err := service.Update(createContext("t", true), &request, &response)

		assertError(err, nil, t)

		if fakeTask1.Title != request.Title {
			t.Errorf("Title hasn't updated: wanted %v got %v", request.Title, fakeTask1.Title)
		}

		if fakeTask1.Description != request.Description {
			t.Errorf("Description hasn't updated: wanted %v got %v", request.Description, fakeTask1.Description)
		}
	})
}

func TestChangeDailyDoStatus(t *testing.T) {

	t.Run("update but repo returns error", func(t *testing.T) {

		service := createService(true, false, true)

		request := taskPb.DailyDoStatusRequest{}
		response := taskPb.Response{}

		err := service.ChangeDailyDoStatus(createContext("t", true), &request, &response)

		assertError(err, errFake, t)

	})

	t.Run("update but returns an error for no token in metadata", func(t *testing.T) {
		service := createService(false, true, true)

		request := taskPb.DailyDoStatusRequest{}
		response := taskPb.Response{}

		err := service.ChangeDailyDoStatus(createContext("", true), &request, &response)

		assertError(err, errFake, t)
	})

	t.Run("update but returns an error for metadata provided", func(t *testing.T) {
		service := createService(false, true, true)

		request := taskPb.DailyDoStatusRequest{}
		response := taskPb.Response{}

		err := service.ChangeDailyDoStatus(createContext("", false), &request, &response)

		assertError(err, errNoMetaData, t)
	})

	t.Run("task not found", func(t *testing.T) {
		service := createService(false, false, true)

		request := taskPb.DailyDoStatusRequest{
			TaskId: "not found",
		}
		response := taskPb.Response{}

		err := service.ChangeDailyDoStatus(createContext("t", true), &request, &response)

		assertError(err, errTaskNotFound, t)
	})

	t.Run("tasks user id doesn't match id in token", func(t *testing.T) {
		service := createService(false, false, false)

		request := taskPb.DailyDoStatusRequest{
			TaskId: "123",
		}
		response := taskPb.Response{}

		err := service.ChangeDailyDoStatus(createContext("t", true), &request, &response)

		assertError(err, errTaskUserIDNotMatched, t)
	})

	t.Run("try to set a task as daily do but a different task is already a daily do", func(t *testing.T) {
		service := createService(false, false, false)

		request := taskPb.DailyDoStatusRequest{
			TaskId: "456",
			Status: true,
		}
		response := taskPb.Response{}

		err := service.ChangeDailyDoStatus(createContext("t", true), &request, &response)

		assertError(err, errDailyDoAlreadyExists, t)
	})

	t.Run("set fakeTask1 for user 1 not as daily do", func(t *testing.T) {

		service := createService(false, false, true)

		request := taskPb.DailyDoStatusRequest{
			TaskId: "123",
			Status: false,
		}

		response := taskPb.Response{}

		err := service.ChangeDailyDoStatus(createContext("t", true), &request, &response)

		assertError(err, nil, t)

		if fakeTask1.DailyDo != request.Status {
			t.Errorf("Daily Do hasn't updated: wanted %v got %v", request.Status, fakeTask1.DailyDo)
		}

	})

	t.Run("set fakeTask1 for user 1 as daily do", func(t *testing.T) {

		service := createService(false, false, true)

		request := taskPb.DailyDoStatusRequest{
			TaskId: "123",
			Status: true,
		}

		response := taskPb.Response{}

		err := service.ChangeDailyDoStatus(createContext("t", true), &request, &response)

		assertError(err, nil, t)

		if fakeTask1.DailyDo != request.Status {
			t.Errorf("Daily Do hasn't updated: wanted %v got %v", request.Status, fakeTask1.DailyDo)
		}
	})
}

func TestCompleteTask(t *testing.T) {
	t.Run("update but repo returns error", func(t *testing.T) {

		service := createService(true, false, true)

		request := taskPb.CompleteTaskRequest{}
		response := taskPb.Response{}

		err := service.CompleteTask(createContext("t", true), &request, &response)

		assertError(err, errFake, t)

	})

	t.Run("update but returns an error for no token in metadata", func(t *testing.T) {
		service := createService(false, true, true)

		request := taskPb.CompleteTaskRequest{}
		response := taskPb.Response{}

		err := service.CompleteTask(createContext("", true), &request, &response)

		assertError(err, errFake, t)
	})

	t.Run("update but returns an error for metadata provided", func(t *testing.T) {
		service := createService(false, true, true)

		request := taskPb.CompleteTaskRequest{}
		response := taskPb.Response{}

		err := service.CompleteTask(createContext("", false), &request, &response)

		assertError(err, errNoMetaData, t)
	})

	t.Run("task not found", func(t *testing.T) {
		service := createService(false, false, true)

		request := taskPb.CompleteTaskRequest{
			TaskId: "not found",
		}
		response := taskPb.Response{}

		err := service.CompleteTask(createContext("t", true), &request, &response)

		assertError(err, errTaskNotFound, t)
	})

	t.Run("tasks user id doesn't match id in token", func(t *testing.T) {
		service := createService(false, false, false)

		request := taskPb.CompleteTaskRequest{
			TaskId: "123",
		}
		response := taskPb.Response{}

		err := service.CompleteTask(createContext("t", true), &request, &response)

		assertError(err, errTaskUserIDNotMatched, t)
	})

	t.Run("set fakeTask4 for user 1 as done", func(t *testing.T) {

		service := createService(false, false, true)

		request := taskPb.CompleteTaskRequest{
			TaskId:    "111",
			Completed: true,
		}

		response := taskPb.Response{}

		err := service.CompleteTask(createContext("t", true), &request, &response)

		assertError(err, nil, t)

		if fakeTask4.CompletedDate == 0 {
			t.Errorf("Daily Do hasn't updated: wanted %v got %v", 0, fakeTask1.CompletedDate)
		}
	})

	t.Run("set fakeTask1 for user 1 as done and unsets daily do status", func(t *testing.T) {

		service := createService(false, false, true)

		request := taskPb.CompleteTaskRequest{
			TaskId:    "123",
			Completed: true,
		}

		response := taskPb.Response{}

		err := service.CompleteTask(createContext("t", true), &request, &response)

		assertError(err, nil, t)

		if fakeTask1.CompletedDate == 0 {
			t.Errorf("Completed Date hasn't updated: wanted %v got %v", 0, fakeTask1.CompletedDate)
		}

		if fakeTask1.DailyDo {
			t.Errorf("Daily do hasn't been updated after being completed. Wanted %v but got %v", false, fakeTask1.DailyDo)
		}
	})

	t.Run("set fakeTask2 for user 1 as not completed", func(t *testing.T) {

		service := createService(false, false, true)

		request := taskPb.CompleteTaskRequest{
			TaskId:    "456",
			Completed: false,
		}

		response := taskPb.Response{}

		err := service.CompleteTask(createContext("t", true), &request, &response)

		assertError(err, nil, t)

		if fakeTask2.CompletedDate != 0 {
			t.Errorf("Completed Date hasn't updated: wanted %v got %v", 0, fakeTask2.CompletedDate)
		}
	})
}
