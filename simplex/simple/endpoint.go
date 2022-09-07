package simplex

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

// The request response struct
type terminologySimplifyConstructRequest struct {
	ID int `json:"id"`
}

type terminologySimplifyResponse struct {
	SimpleTerms SimpleTerms `json:"simpleterms"`
	Err         string      `json:"err,omitempty"`
}

// terminologySimplifyEndpoint: This function calls the TerminologySimplified method of the TerminologyService
// It creates the endpoint for the method that is exposed by the transport.
func terminologySimplifyEndpoint(svc TerminologyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(terminologySimplifyConstructRequest)
		v, err := svc.TerminologySimplified(ctx, req.ID)
		if err != nil {
			return terminologySimplifyResponse{v, err.Error()}, nil
		}
		return terminologySimplifyResponse{v, ""}, nil
	}
}
