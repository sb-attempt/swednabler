package curat

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"
)

// NewHttpServer: This is to expose the endpoints.
// The endpoints are to generate jwt token and also to validate them.
func NewHttpServer(svc TerminologyService, logger kitlog.Logger) *mux.Router {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeErrorResponse),
		kithttp.ServerFinalizer(newServerFinalizer(logger)),
	}
	terminologyListHandler := kithttp.NewServer(
		terminologyListEndpoint(svc),
		decodeTerminologyListRequest,
		encodeResponse,
		options...,
	)
	terminologySimplifyHandler := kithttp.NewServer(
		terminologySimplifyEndpoint(svc),
		decodeTerminologySimplifyRequest,
		encodeResponse,
		options...,
	)

	r := mux.NewRouter()
	r.Methods("GET").Path("/v1/term/list").Handler(terminologyListHandler)
	r.Methods("POST").Path("/v1/term/simplify").Handler(terminologySimplifyHandler)
	return r
}

// The below methods are standard encode/decode functions for the incoming/outgoing requests/responses.
func decodeTerminologyListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return struct{}{}, errors.New("authorization header missing")
	}
	var request = terminologyListRequest{
		Token: token,
	}
	return request, nil
}

func decodeTerminologySimplifyRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return struct{}{}, errors.New("authorization header missing")
	}
	var reqReceived terminologySimplifyRequest
	if err := json.NewDecoder(r.Body).Decode(&reqReceived); err != nil {
		return nil, err
	}

	var constructRequest = terminologySimplifyConstructRequest{
		ID:    reqReceived.ID,
		Token: token,
	}
	return constructRequest, nil
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
