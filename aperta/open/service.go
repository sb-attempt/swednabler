package open

import (
	"context"
	"errors"
)

// Go-kit: Combine all related function into an interface
// The Service interface defines the structure of both authenticate and authorize methods.
type Service interface {
	Authenticate(ctx context.Context, jwt JwtService, username, password string) (string, error)
	Authorize(ctx context.Context, jwt JwtService, token string) (bool, int)
}

// The constructor for service
func NewService() *service {
	return &service{}
}

// The struct that implements the interface.
type service struct{}

// Creating a user map; this is a very crude of defining the identity.
// It solves the current purpose to keep this code experimental.
// TODO: move this map tp a dataset?
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Authenticate: This function basically checks if the credential provided match the one in the map.
// If authentication is successful then the token is generated and returned to the calling function.
func (s service) Authenticate(ctx context.Context, jwt JwtService, username, password string) (string, error) {
	passwordExpected, ok := users[username]

	if !ok || password != passwordExpected {
		return "", errors.New("the user is invalid")
	}
	token, err := jwt.GetToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Authorize: A very simple function that just authorize the incoming token
// It return a bool and an errorCode that tells tha caller if token is valid or has expired.
// The call can authorize based on results
func (s service) Authorize(ctx context.Context, jwt JwtService, token string) (bool, int) {
	valid, errorCode := jwt.ValidateToken(token)
	return valid, errorCode
}
