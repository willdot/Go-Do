package main

import (
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"

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

	t.Run("token password no longer valid", func(t *testing.T) {
		service := createService(false)

		hashedPass, _ := bcrypt.GenerateFromPassword([]byte("new"), bcrypt.DefaultCost)
		user := pb.User{
			Id:       "123",
			Name:     "Fake",
			Email:    "fake@fake.com",
			Password: string(hashedPass),
			Company:  "fake",
		}

		token, _ := service.tokenService.Encode(&user)
		request := pb.Token{Token: token}
		response := pb.Token{}

		got := service.ValidateToken(createContext(), &request, &response)

		if got != errTokenPasswordNotValid {
			t.Errorf("wanted %v but got %v", nil, got)
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
