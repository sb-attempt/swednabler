package auth

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

// Structs that defines the structure of incoming request/response for both authenticate/authorize requests
type authenticateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authenticateResponse struct {
	Token string `json:"token,omitempty"`
	Err   string `json:"err,omitempty"`
}

type authorizeRequest struct {
	Token string `json:"token"`
}

type authorizeResponse struct {
	Valid   bool `json:"valid"`
	ErrCode int  `json:"err,omitempty"`
}

// authenticateEndpoint: This function accepts the Service and JwtService as input.
// It checks if the incoming credentials are valid and return a JWT token with a set expiry
// Based on go-kit utility this adapter converts the  method to endpoint and then the transport exposes it.
// https://github.com/go-kit/kit
func authenticateEndpoint(svc Service, jwt JwtService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(authenticateRequest)
		// Authenticate method that returns the token
		v, err := svc.Authenticate(ctx, jwt, req.Username, req.Password)
		if err != nil {
			return authenticateResponse{v, err.Error()}, nil
		}
		return authenticateResponse{v, ""}, nil
	}
}

// authorizeEndpoint: This function accepts the Service and JwtService as input.
// It check if the token passed is valid and returns a bool value based on result.
// The adapter to convert method to endpoint.
func authorizeEndpoint(svc Service, jwt JwtService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(authorizeRequest)
		// Authorize method that validates the incoming token.
		v, errCode := svc.Authorize(ctx, jwt, req.Token)
		return authorizeResponse{v, errCode}, nil
	}
}
