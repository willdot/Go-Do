package main

import (
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
}

// UserRepository is a datastore
type UserRepository struct {
	Session *gocql.Session
}
