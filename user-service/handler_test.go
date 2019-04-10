package main

import (
	"errors"
	"net/http"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"

	pb "github.com/willdot/go-do/user-service/proto/auth"
)

func TestCreate(t *testing.T) {

	t.Run("returns a user", func(t *testing.T) {
		service := createService(false)

		user := pb.User{
			Name:     "Fake",
			Email:    "fake@fake.com",
			Password: "fake",
			Company:  "fake",
		}

		response := pb.Response{}

		got := service.Create(createContext(), &user, &response)

		if got != nil {
			t.Errorf("wanted %v but got %v", nil, got)
		}
	})

	t.Run("returns an error", func(t *testing.T) {
		service := createService(true)

		user := pb.User{
			Name:     "Fake",
			Email:    "fake@fake.com",
			Password: "fake",
			Company:  "fake",
		}

		response := pb.Response{}

		got := service.Create(createContext(), &user, &response)

		if got != errFake {
			t.Errorf("wanted %v but got %v", nil, got)
		}
	})

}

func TestGet(t *testing.T) {

	t.Run("returns a user", func(t *testing.T) {
		service := createService(false)

		user := pb.User{
			Name:     "Fake",
			Email:    "fake@fake.com",
			Password: "fake",
			Company:  "fake",
		}

		response := pb.Response{}

		got := service.Get(createContext(), &user, &response)

		if got != nil {
			t.Errorf("wanted %v but got %v", nil, got)
		}
	})

	t.Run("returns an error", func(t *testing.T) {
		service := createService(true)

		user := pb.User{
			Name:     "Fake",
			Email:    "fake@fake.com",
			Password: "fake",
			Company:  "fake",
		}

		response := pb.Response{}

		got := service.Get(createContext(), &user, &response)

		if got != errFake {
			t.Errorf("wanted %v but got %v", nil, got)
		}
	})

}

func TestGetAll(t *testing.T) {

	t.Run("returns a user", func(t *testing.T) {
		service := createService(false)

		request := pb.Request{}

		response := pb.Response{}

		got := service.GetAll(createContext(), &request, &response)

		if got != nil {
			t.Errorf("wanted %v but got %v", nil, got)
		}
	})

	t.Run("returns an error", func(t *testing.T) {
		service := createService(true)

		request := pb.Request{}

		response := pb.Response{}

		got := service.GetAll(createContext(), &request, &response)

		if got != errFake {
			t.Errorf("wanted %v but got %v", nil, got)
		}
	})

}

func TestUpdate(t *testing.T) {

	t.Run("returns a user", func(t *testing.T) {
		service := createService(false)

		user := pb.User{
			Name:     "Fake",
			Email:    "fake@fake.com",
			Password: "fake",
			Company:  "fake",
		}
		response := pb.Response{}

		got := service.Update(createContext(), &user, &response)

		if got != nil {
			t.Errorf("wanted %v but got %v", nil, got)
		}
	})

	t.Run("returns an error", func(t *testing.T) {
		service := createService(true)

		user := pb.User{
			Name:     "Fake",
			Email:    "fake@fake.com",
			Password: "fake",
			Company:  "fake",
		}
		response := pb.Response{}

		got := service.Update(createContext(), &user, &response)

		if got != errFake {
			t.Errorf("wanted %v but got %v", nil, got)
		}
	})

}

func TestValidateToken(t *testing.T) {

	t.Run("token is valid", func(t *testing.T) {
		service := createService(false)

		user := &fakeUser
		user.Id = "123"

		token, _ := service.tokenService.Encode(user)
		request := pb.Token{Token: token}
		response := pb.Token{}

		got := service.ValidateToken(createContext(), &request, &response)

		if got != nil {
			t.Errorf("wanted %v but got %v", nil, got)
		}
	})

	t.Run("token has expired", func(t *testing.T) {
		service := createService(true)

		user := &fakeUser
		user.Id = "123"

		token, _ := service.tokenService.Encode(user)
		request := pb.Token{Token: token}
		response := pb.Token{}

		got := service.ValidateToken(createContext(), &request, &response)

		if !strings.Contains(got.Error(), "token is expired by") {
			t.Errorf("wanted 'token is expired by......' but got %v", got)
		}
	})

	t.Run("user not known", func(t *testing.T) {
		service := createService(false)

		user := &fakeUser
		user.Id = ""

		token, _ := service.tokenService.Encode(user)
		request := pb.Token{Token: token}
		response := pb.Token{}

		got := service.ValidateToken(createContext(), &request, &response)

		if got != errUnknownUser {
			t.Errorf("wanted %v but got %v", nil, got)
		}
	})

}

func TestAuth(t *testing.T) {

	t.Run("auth is fine", func(t *testing.T) {
		service := createService(false)

		response := pb.Token{}

		user := pb.User{
			Password: "test",
		}

		got := service.Auth(createContext(), &user, &response)

		if got != nil {
			t.Errorf("wanted %v but got %v", nil, got)
		}
	})

	t.Run("user not found", func(t *testing.T) {
		service := createService(true)

		response := pb.Token{}

		got := service.Auth(createContext(), &fakeUser, &response)

		if got != errFake {
			t.Errorf("wanted %v but got %v", nil, got)
		}
	})

	t.Run("password doesn't match", func(t *testing.T) {
		service := createService(false)

		response := pb.Token{}

		user := pb.User{
			Password: "wrong",
		}

		got := service.Auth(createContext(), &user, &response)

		if !strings.Contains(got.Error(), "hashedPassword is not the hash of the given password") {
			t.Errorf("wanted 'hashedPassword is not the hash of the given password......' but got %v", got)
		}
	})
}
func TestPasswordChange(t *testing.T) {

	t.Run("password changed successfully", func(t *testing.T) {
		service := createService(false)

		response := pb.Token{}

		request := pb.PasswordChange{
			Email:       "fake@fake.com",
			OldPassword: "test",
			NewPassword: "new",
		}

		got := service.ChangePassword(createContext(), &request, &response)

		if got != nil {
			t.Errorf("wanted %v but got %v", nil, got)
		}
	})

	t.Run("user not found", func(t *testing.T) {
		service := createService(true)

		response := pb.Token{}

		request := pb.PasswordChange{
			Email:       "notreal",
			OldPassword: "test",
			NewPassword: "new",
		}

		got := service.ChangePassword(createContext(), &request, &response)

		if got != errFake {
			t.Errorf("wanted %v but got %v", nil, got)
		}
	})

	t.Run("old password incorrect", func(t *testing.T) {
		service := createService(false)

		response := pb.Token{}

		request := pb.PasswordChange{
			Email:       "notreal",
			OldPassword: "wrong",
			NewPassword: "new",
		}

		got := service.ChangePassword(createContext(), &request, &response)

		if !strings.Contains(got.Error(), "hashedPassword is not the hash of the given password") {
			t.Errorf("wanted 'hashedPassword is not the hash of the given password......' but got %v", got)
		}
	})

	

}

type fakeRepo struct {
	returnError bool
	users       []*pb.User
}

var errFake = errors.New("This is a fake error message")

func (f *fakeRepo) GetAll() ([]*pb.User, error) {

	if f.returnError {
		return nil, errFake
	}
	return f.users, nil
}

func (f *fakeRepo) Get(id string) (*pb.User, error) {

	if f.returnError {
		return nil, errFake
	}
	return nil, nil
}

func (f *fakeRepo) Create(user *pb.User) error {

	if f.returnError {
		return errFake
	}
	return nil
}

func (f *fakeRepo) GetByEmail(email string) (*pb.User, error) {

	if f.returnError {
		return nil, errFake
	}

	user := f.users[0]
	return user, nil
}

func (f *fakeRepo) Update(user *pb.User) error {

	if f.returnError {
		return errFake
	}
	return nil
}

func (f *fakeRepo) UpdatePassword(id, password string) error {

	if f.returnError {
		return errFake
	}
	return nil
}

var fakeUser = pb.User{
	Name:     "Fake",
	Email:    "fake@fake.com",
	Password: "$2a$10$cSOEkdxPPOrX8h/t3/Aw5e.vludnAzMGU38I3Cv0V/GAAwaqyJDaK",
	Company:  "fake",
}

func createService(returnError bool) userHandler {

	var users []*pb.User

	users = append(users, &fakeUser)

	fakeRepo := &fakeRepo{returnError, users}

	expireTime := time.Now().Add(time.Hour * 1).Unix()

	if returnError {
		expireTime = time.Now().AddDate(-1, 0, 0).Unix()
	}

	tokenService := TokenService{fakeRepo, expireTime}

	service := userHandler{fakeRepo, tokenService}

	return service
}

func createContext() context.Context {
	req, _ := http.NewRequest(http.MethodPost, "/", nil)

	ctx, _ := context.WithDeadline(req.Context(), time.Now())

	return ctx
}
