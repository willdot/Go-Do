package main

import (
	"errors"
	"strconv"
	"time"

	"github.com/gocql/gocql"
	taskPb "github.com/willdot/go-do/task-service/proto/task"
)

var errTaskNotFound = errors.New("Task not found")
var errTaskUserIDNotMatched = errors.New("The user id for the task provided doesn't match user id from token")

// Repository ..
type Repository interface {
	Get(string) ([]*taskPb.Task, error)
	Create(*taskPb.Task) error
	Update(*taskPb.Task) error
	SetDailyDoStatus(*taskPb.Task) error
	GetDailyDoForUser(string) (*taskPb.Task, error)
	CompleteTask(*taskPb.Task) error
}

// TaskRepository is a datastore
type TaskRepository struct {
	Session *gocql.Session
}

// Get will get tasks for a user
func (repo *TaskRepository) Get(userID string) ([]*taskPb.Task, error) {
	var tasks []*taskPb.Task

	m := map[string]interface{}{}

	query := repo.Session.Query("SELECT * FROM task WHERE userId =?", userID)
	iterable := query.Iter()

	for iterable.MapScan(m) {

		tasks = append(tasks, &taskPb.Task{
			Id:            m["id"].(gocql.UUID).String(),
			Title:         m["title"].(string),
			Description:   m["description"].(string),
			UserId:        m["userid"].(string),
			DailyDo:       m["dailydo"].(bool),
			CompletedDate: m["completeddate"].(time.Time).Unix(),
			CreatedDate:   m["createddate"].(time.Time).Unix(),
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

	query := repo.Session.Query("SELECT * FROM task WHERE id = ?", task.Id)
	iterable := query.Consistency(gocql.One).Iter()

	for iterable.MapScan(m) {
		found = true
		existingTask = taskPb.Task{
			UserId: m["userid"].(string),
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
func (repo *TaskRepository) GetDailyDoForUser(userID string) (*taskPb.Task, error) {

	var dailyDo *taskPb.Task
	m := map[string]interface{}{}

	query := repo.Session.Query("SELECT * FROM task WHERE dailyDo = true and userId =? ALLOW FILTERING", userID)
	iterable := query.Consistency(gocql.One).Iter()

	for iterable.MapScan(m) {
		dailyDo = &taskPb.Task{
			Id:            m["id"].(gocql.UUID).String(),
			Title:         m["title"].(string),
			Description:   m["description"].(string),
			UserId:        m["userid"].(string),
			CreatedDate:   m["createddate"].(time.Time).Unix(),
			CompletedDate: m["completeddate"].(time.Time).Unix(),
			DailyDo:       m["dailydo"].(bool),
		}
	}

	return dailyDo, nil
}

// CompleteTask sets the completed date time of the task. Sets to 0 if it's being un completed
func (repo *TaskRepository) CompleteTask(task *taskPb.Task) error {

	var found = false

	var existingTask taskPb.Task
	m := map[string]interface{}{}

	query := repo.Session.Query("SELECT * FROM task WHERE taskId=? LIMIT 1", task.Id)
	iterable := query.Consistency(gocql.One).Iter()

	for iterable.MapScan(m) {
		found = true
		existingTask = taskPb.Task{
			UserId:  m["userid"].(string),
			DailyDo: m["dailydo"].(bool),
		}
	}

	if !found {
		return errTaskNotFound
	}

	if existingTask.UserId != task.UserId {
		return errTaskUserIDNotMatched
	}

	parameters := []string{strconv.FormatInt(task.CompletedDate, 10)}
	queryString := "UPDATE task SET completedDate =?"

	if existingTask.DailyDo != task.DailyDo {
		queryString += ", dailyDo =?"
		parameters = append(parameters, strconv.FormatBool(task.DailyDo))
	}

	parameters = append(parameters, task.Id)

	queryString += " where id = ?"

	err := repo.Session.Query(queryString, parameters).Exec()

	if err != nil {
		return err
	}

	return err

}
