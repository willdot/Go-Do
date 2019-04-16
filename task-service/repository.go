package main

import (
	"errors"

	"github.com/gocql/gocql"
	taskPb "github.com/willdot/go-do/task-service/proto/task"
)

var errTaskNotFound = errors.New("Task not found")
var errTaskUserIDNotMatched = errors.New("The user id for the task provided doesn't match user id from token")

// Repository ..
type Repository interface {
	Get(userID string) ([]*taskPb.Task, error)
	Create(*taskPb.Task) error
	Update(*taskPb.Task) error
	SetDailyDoStatus(*taskPb.Task) error
	GetDailyDoForUser(string) (*taskPb.Task, error)
}

// TaskRepository is a datastore
type TaskRepository struct {
	Session *gocql.Session
}

// Get will get tasks for a user
func (repo *TaskRepository) Get(userID string) ([]*taskPb.Task, error) {
	var tasks []*taskPb.Task

	m := map[string]interface{}{}

	query := repo.Session.Query("SELECT * FROM task WHERE userId =?")
	iterable := query.Iter()

	for iterable.MapScan(m) {
		tasks = append(tasks, &taskPb.Task{
			Id:            m["id"].(gocql.UUID).String(),
			Title:         m["title"].(string),
			Description:   m["description"].(string),
			UserId:        m["userId"].(string),
			CreatedDate:   m["createdDate"].(int64),
			CompletedDate: m["completedDate"].(int64),
			DailyDo:       m["dailyDo"].(bool),
		})

		m = map[string]interface{}{}
	}

	return tasks, nil
}

// Create will create a new task
func (repo *TaskRepository) Create(task *taskPb.Task) error {
	gocqlUUID := gocql.TimeUUID()

	err := repo.Session.Query(`
	INSERT INTO task (id, title, description, userId, createdDate, dailyDo) VALUES (?,?,?,?,?,?)`,
		gocqlUUID, task.Title, task.Description, task.UserId, task.CreatedDate, task.DailyDo).Exec()

	task.Id = gocqlUUID.String()

	return err
}

// Update will update a task
func (repo *TaskRepository) Update(task *taskPb.Task) error {

	var found = false

	var existingTask taskPb.Task
	m := map[string]interface{}{}

	query := repo.Session.Query("SELECT * FROM task WHERE taskId=? LIMIT 1", task.Id)
	iterable := query.Consistency(gocql.One).Iter()

	for iterable.MapScan(m) {
		found = true
		existingTask = taskPb.Task{
			UserId: m["userId"].(string),
		}
	}

	if !found {
		return errTaskNotFound
	}

	if existingTask.UserId != task.UserId {
		return errTaskUserIDNotMatched
	}

	if task.Title == "" {
		task.Title = existingTask.Title
	}

	if task.Description == "" {
		task.Description = existingTask.Description
	}

	err := repo.Session.Query("UPDATE task SET title =?, description = ? where id = ?", task.Title, task.Description, task.Id).Exec()

	return err
}

// SetDailyDoStatus will set a task as a daily do
func (repo *TaskRepository) SetDailyDoStatus(task *taskPb.Task) error {

	var found = false

	var existingTask taskPb.Task
	m := map[string]interface{}{}

	query := repo.Session.Query("SELECT * FROM task WHERE taskId=? LIMIT 1", task.Id)
	iterable := query.Consistency(gocql.One).Iter()

	for iterable.MapScan(m) {
		found = true
		existingTask = taskPb.Task{
			UserId: m["userId"].(string),
		}
	}

	if !found {
		return errTaskNotFound
	}

	if existingTask.UserId != task.UserId {
		return errTaskUserIDNotMatched
	}

	err := repo.Session.Query("UPDATE task SET dailyDo =? where id = ?", task.DailyDo, task.Id).Exec()

	return err
}

// GetDailyDoForUser will get a daily do task for a user
func (repo *TaskRepository) GetDailyDoForUser(string) (*taskPb.Task, error) {

	var dailyDo *taskPb.Task
	m := map[string]interface{}{}

	query := repo.Session.Query("SELECT * FROM task WHERE dailyDo=true LIMIT 1")
	iterable := query.Consistency(gocql.One).Iter()

	for iterable.MapScan(m) {
		dailyDo = &taskPb.Task{
			Id:            m["id"].(gocql.UUID).String(),
			Title:         m["title"].(string),
			Description:   m["description"].(string),
			UserId:        m["userId"].(string),
			CreatedDate:   m["createdDate"].(int64),
			CompletedDate: m["completedDate"].(int64),
			DailyDo:       m["dailyDo"].(bool),
		}
	}

	return dailyDo, nil
}
