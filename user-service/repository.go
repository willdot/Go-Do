package main

import (
	"errors"

	"github.com/gocql/gocql"
	pb "github.com/willdot/go-do/user-service/proto/auth"
)

// Repository ..
type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmail(email string) (*pb.User, error)
	Update(user *pb.User) error
	UpdatePassword(id, password string) error
}

// UserRepository is a datastore
type UserRepository struct {
	Session *gocql.Session
}

// GetAll will get all users from database
func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User

	m := map[string]interface{}{}

	query := repo.Session.Query("SELECT * FROM user")
	iterable := query.Iter()

	for iterable.MapScan(m) {
		users = append(users, &pb.User{
			Id:       m["id"].(gocql.UUID).String(),
			Name:     m["name"].(string),
			Email:    m["email"].(string),
			Password: m["password"].(string),
			Company:  m["company"].(string),
		})
		m = map[string]interface{}{}
	}

	return users, nil
}

// Get will get a single user
func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var found = false
	var user pb.User
	m := map[string]interface{}{}

	query := repo.Session.Query("SELECT * FROM user WHERE id=? LIMIT 1", id)
	iterable := query.Consistency(gocql.One).Iter()

	for iterable.MapScan(m) {
		found = true
		user = pb.User{
			Id:       m["id"].(gocql.UUID).String(),
			Name:     m["name"].(string),
			Email:    m["email"].(string),
			Password: m["password"].(string),
			Company:  m["company"].(string),
		}
	}

	if found {
		return &user, nil
	}

	return nil, errors.New("User can't be found")
}

// GetByEmail will get a user by email
func (repo *UserRepository) GetByEmail(email string) (*pb.User, error) {

	var found = false
	var user pb.User
	m := map[string]interface{}{}

	query := repo.Session.Query("SELECT * FROM user WHERE email=? LIMIT 1", email)
	iterable := query.Consistency(gocql.One).Iter()

	for iterable.MapScan(m) {
		found = true
		user = pb.User{
			Id:       m["id"].(gocql.UUID).String(),
			Name:     m["name"].(string),
			Email:    m["email"].(string),
			Password: m["password"].(string),
			Company:  m["company"].(string),
		}
	}

	if found {
		return &user, nil
	}

	return nil, errors.New("User can't be found")
}

// Create will create a new user
func (repo *UserRepository) Create(user *pb.User) error {

	gocqlUUID := gocql.TimeUUID()

	err := repo.Session.Query(`
	INSERT INTO user (id, name, email, password, company) VALUES (?,?,?,?,?)`,
		gocqlUUID, user.Name, user.Email, user.Password, user.Company).Exec()

	return err
}

// Update will update a user
func (repo *UserRepository) Update(user *pb.User) error {

	err := repo.Session.Query(`UPDATE user SET name = ?, company = ? where id = ?`, user.Name, user.Company, user.Id).Exec()

	return err
}

// UpdatePassword updates the users password
func (repo *UserRepository) UpdatePassword(id, password string) error {

	err := repo.Session.Query(`UPDATE user SET password = ? where id = ?`, password, id).Exec()

	return err
}