package open

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"
)

// NewHttpServer: This is to expose the endpoints.
// The endpoints are to generate jwt token and also to validate them.
func NewHttpServer(svc Service, jwt JwtService, logger kitlog.Logger) *mux.Router {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeErrorResponse),
		kithttp.ServerFinalizer(newServerFinalizer(logger)),
	}
	authenticateHandler := kithttp.NewServer(
		authenticateEndpoint(svc, jwt),
		decodeAuthenticateRequest,
		encodeResponse,
		options...,
	)
	authorizeHandler := kithttp.NewServer(
		authorizeEndpoint(svc, jwt),
		decodeAuthorizeRequest,
		encodeResponse,
		options...,
	)
	// The basic mux router
	r := mux.NewRouter()
	r.Methods("POST").Path("/v1/token").Handler(authenticateHandler)
	r.Methods("POST").Path("/v1/token/validate").Handler(authorizeHandler)
	return r
}

// The below methods are standard encode/decode functions for the incoming/outgoing requests/responses.
func decodeAuthenticateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request authenticateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeAuthorizeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request authorizeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func newServerFinalizer(logger kitlog.Logger) kithttp.ServerFinalizerFunc {
	return func(ctx context.Context, code int, r *http.Request) {
		_ = logger.Log("status", code, "path", r.RequestURI, "method", r.Method)
	}
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
