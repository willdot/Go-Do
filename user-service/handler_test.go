package main

import (
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"

	authPb "github.com/willdot/go-do/user-service/proto/auth"
)

func assertError(got, want error, t *testing.T) {
	if got != want {
		t.Errorf("got error '%v' but want error '%v'", got, want)
	}
}

func TestCreate(t *testing.T) {

	t.Run("returns a user", func(t *testing.T) {
		service := createService(false)

		user := authPb.User{
			Name:     "Fake",
			Email:    "fake@fake.com",
			Password: "fake",
			Company:  "fake",
		}

		response := authPb.Response{}

		err := service.Create(createContext(), &user, &response)

		assertError(err, nil, t)
	})

	t.Run("returns an error", func(t *testing.T) {
		service := createService(true)

		user := authPb.User{
			Name:     "Fake",
			Email:    "fake@fake.com",
			Password: "fake",
			Company:  "fake",
		}

		response := authPb.Response{}

		err := service.Create(createContext(), &user, &response)

		assertError(err, errFake, t)
	})

}

func TestGet(t *testing.T) {

	t.Run("returns a user", func(t *testing.T) {
		service := createService(false)

		user := authPb.User{
			Name:     "Fake",
			Email:    "fake@fake.com",
			Password: "fake",
			Company:  "fake",
		}

		response := authPb.Response{}

		err := service.Get(createContext(), &user, &response)

		assertError(err, nil, t)
	})

	t.Run("returns an error", func(t *testing.T) {
		service := createService(true)

		user := authPb.User{
			Name:     "Fake",
			Email:    "fake@fake.com",
			Password: "fake",
			Company:  "fake",
		}

		response := authPb.Response{}

		err := service.Get(createContext(), &user, &response)

		assertError(err, errFake, t)
	})

}

func TestGetAll(t *testing.T) {

	t.Run("returns a user", func(t *testing.T) {
		service := createService(false)

		request := authPb.Request{}

		response := authPb.Response{}

		err := service.GetAll(createContext(), &request, &response)

		assertError(err, nil, t)
	})

	t.Run("returns an error", func(t *testing.T) {
		service := createService(true)

		request := authPb.Request{}

		response := authPb.Response{}

		err := service.GetAll(createContext(), &request, &response)

		assertError(err, errFake, t)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("returns a user", func(t *testing.T) {
		service := createService(false)

		user := authPb.User{
			Name:     "Fake",
			Email:    "fake@fake.com",
			Password: "fake",
			Company:  "fake",
		}
		response := authPb.Response{}

		err := service.Update(createContext(), &user, &response)

		assertError(err, nil, t)
	})

	t.Run("returns an error", func(t *testing.T) {
		service := createService(true)

		user := authPb.User{
			Name:     "Fake",
			Email:    "fake@fake.com",
			Password: "fake",
			Company:  "fake",
		}
		response := authPb.Response{}

		err := service.Update(createContext(), &user, &response)

		assertError(err, errFake, t)
	})

}

func TestValidateToken(t *testing.T) {

	t.Run("token is valid", func(t *testing.T) {
		service := createService(false)

		user := &fakeUser
		user.Id = "123"

		token, _ := service.tokenService.Encode(user)
		request := authPb.Token{Token: token}
		response := authPb.Token{}

		err := service.ValidateToken(createContext(), &request, &response)

		assertError(err, nil, t)
	})

	t.Run("token has expired", func(t *testing.T) {
		service := createService(true)

		user := &fakeUser
		user.Id = "123"

		token, _ := service.tokenService.Encode(user)
		request := authPb.Token{Token: token}
		response := authPb.Token{}

		err := service.ValidateToken(createContext(), &request, &response)

		if !strings.Contains(err.Error(), "token is expired by") {
			t.Errorf("wanted 'token is expired by......' but got %v", err)
		}
	})

	t.Run("token password no longer valid", func(t *testing.T) {
		service := createService(false)

		hashedPass, _ := bcrypt.GenerateFromPassword([]byte("new"), bcrypt.DefaultCost)
		user := authPb.User{
			Id:       "123",
			Name:     "Fake",
			Email:    "fake@fake.com",
			Password: string(hashedPass),
			Company:  "fake",
		}

		token, _ := service.tokenService.Encode(&user)
		request := authPb.Token{Token: token}
		response := authPb.Token{}

		err := service.ValidateToken(createContext(), &request, &response)

		assertError(err, errTokenPasswordNotValid, t)
	})

	t.Run("user not known", func(t *testing.T) {
		service := createService(false)

		user := &fakeUser
		user.Id = ""

		token, _ := service.tokenService.Encode(user)
		request := authPb.Token{Token: token}
		response := authPb.Token{}

		err := service.ValidateToken(createContext(), &request, &response)

		assertError(err, errUnknownUser, t)
	})

}

func TestAuth(t *testing.T) {

	t.Run("auth is fine", func(t *testing.T) {
		service := createService(false)

		response := authPb.Token{}

		user := authPb.User{
			Password: "test",
		}

		err := service.Auth(createContext(), &user, &response)

		assertError(err, nil, t)
	})

	t.Run("user not found", func(t *testing.T) {
		service := createService(true)

		response := authPb.Token{}

		err := service.Auth(createContext(), &fakeUser, &response)

		assertError(err, errFake, t)
	})

	t.Run("password doesn't match", func(t *testing.T) {
		service := createService(false)

		response := authPb.Token{}

		user := authPb.User{
			Password: "wrong",
		}

		err := service.Auth(createContext(), &user, &response)

		if !strings.Contains(err.Error(), "hashedPassword is not the hash of the given password") {
			t.Errorf("wanted 'hashedPassword is not the hash of the given password......' but got %v", err)
		}
	})
}
func TestPasswordChange(t *testing.T) {

	t.Run("password changed successfully", func(t *testing.T) {
		service := createService(false)

		response := authPb.Token{}

		request := authPb.PasswordChange{
			Email:       "fake@fake.com",
			OldPassword: "test",
			NewPassword: "new",
		}

		err := service.ChangePassword(createContext(), &request, &response)

		assertError(err, nil, t)
	})

	t.Run("user not found", func(t *testing.T) {
		service := createService(true)

		response := authPb.Token{}

		request := authPb.PasswordChange{
			Email:       "notreal",
			OldPassword: "test",
			NewPassword: "new",
		}

		err := service.ChangePassword(createContext(), &request, &response)

		assertError(err, errFake, t)
	})

	t.Run("old password incorrect", func(t *testing.T) {
		service := createService(false)

		response := authPb.Token{}

		request := authPb.PasswordChange{
			Email:       "notreal",
			OldPassword: "wrong",
			NewPassword: "new",
		}

		err := service.ChangePassword(createContext(), &request, &response)

		if !strings.Contains(err.Error(), "hashedPassword is not the hash of the given password") {
			t.Errorf("wanted 'hashedPassword is not the hash of the given password......' but got %v", err)
		}
	})

}
