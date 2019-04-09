package main

import (
	"github.com/dgrijalva/jwt-go"
	pb "github.com/willdot/go-do/user-service/proto/auth"
)

var (
	key = []byte("mysupersecretkey")
)

// CustomClaims ..
type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

// Authable ..
type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

// TokenService ..
type TokenService struct {
	repo       Repository
	expireTime int64
}

// Decode a token
func (s *TokenService) Decode(tokenString string) (*CustomClaims, error) {

	tokenType, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims, nil
	}

	return nil, err
}

// Encode a claim into a JWT
func (s *TokenService) Encode(user *pb.User) (string, error) {
	expireToken := s.expireTime //time.Now().Add(time.Hour * 72).Unix()

	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "go-do.user",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}
