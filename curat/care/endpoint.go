package curat

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

// Structs that defines the structure of incoming request and response for both authenticate and authorize requests
type terminologyListResponse struct {
	Terms Terms  `json:"terms,omitempty"`
	Err   string `json:"err,omitempty"`
}

type terminologyListRequest struct {
	Token string `json:"token"`
}

type terminologySimplifyRequest struct {
	ID int `json:"id"`
}

type terminologySimplifyConstructRequest struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

type terminologySimplifyResponse struct {
	SimpleTerms SimpleTerms `json:"simpleterms,omitempty"`
	Err         string      `json:"err,omitempty"`
}

// terminologyListEndpoint: This function fetch the terminology list from the database.
// It takes the the token and validates it.
// If the token is valid then it returns the list of all the terms along with their IDs in the JSOn format
func terminologyListEndpoint(svc TerminologyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(terminologyListRequest)
		// Fetch the term list from dataset
		v, err := svc.TerminologyList(ctx, req.Token)
		if err != nil {
			return terminologyListResponse{v, err.Error()}, nil
		}
		return terminologyListResponse{v, ""}, nil
	}
}

// terminologySimplifyEndpoint: This function calls the simplify endpoint of the simplex service.
// It accepts the token and ID as input and if token is valid it fetch the data corresponding to the ID.
func terminologySimplifyEndpoint(svc TerminologyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(terminologySimplifyConstructRequest)
		v, err := svc.TerminologySimplified(ctx, req.Token, req.ID)
		if err != nil {
			return terminologySimplifyResponse{v.SimpleTerms, err.Error()}, nil
		}
		return terminologySimplifyResponse{v.SimpleTerms, ""}, nil
	}
}
