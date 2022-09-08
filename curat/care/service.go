package curat

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// Request and response structs
type authorizeRequest struct {
	Token string `json:"token"`
}

type authorizeResponse struct {
	Valid   bool `json:"valid"`
	ErrCode int  `json:"err,omitempty"`
}

// Interface
type TerminologyService interface {
	TerminologyList(ctx context.Context, token string) (Terms, error)
	TerminologySimplified(ctx context.Context, token string, id int) (terminologySimplifyResponse, error)
	ValidateToken(token string) error
}

func NewTerminologyService() *terminologyService {
	return &terminologyService{}
}

// Struct that implements the TerminologyService interface
type terminologyService struct{}

// TerminologyList: This function generates the terminology list if the provided token is valid
func (t terminologyService) TerminologyList(ctx context.Context, token string) (Terms, error) {
	err := t.ValidateToken(token)
	if err != nil {
		return nil, err
	}
	return Data.Glossary, nil
}

// ValidateToken: The common function to validate token.
// The token is validated against Aperta service by calling the validate endpoint with provided token.
// If token is expired or invalid then endpoint denies the request.
func (t terminologyService) ValidateToken(token string) error {
	//Verify if the token is valid
	if ApertaServiceHost == "" {
		ApertaServiceHost = "localhost"
	}
	jwtString := strings.Split(token, "Bearer ")[1]
	requestBody := authorizeRequest{
		Token: jwtString,
	}
	parseJSON, err := json.Marshal(requestBody)
	if err != nil {
		return errors.New("could not validate token internal server error")
	}
	body := bytes.NewBuffer(parseJSON)
	// Call the Aperta service validate endpoint
	resp, err := http.Post("http://"+ApertaServiceHost+":8081/v1/token/validate", "application/json", body)
	if err != nil {
		return errors.New("could not validate token internal server error")
	}
	var result authorizeResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}
	if !result.Valid {
		return errors.New("token is either expired or unauthorized")
	}
	return nil
}

// TerminologySimplified: This function takes an ID and give more information for that particular ID provided.
// It call the simplex service and pass the id as payload, the service returns the simplified response.
func (t terminologyService) TerminologySimplified(ctx context.Context, token string, id int) (terminologySimplifyResponse, error) {
	var simpleTerms terminologySimplifyResponse
	if SimplexServiceHost == "" {
		SimplexServiceHost = "localhost"
	}
	// Validate the token
	err := t.ValidateToken(token)
	if err != nil {
		return simpleTerms, err
	}

	requestBody := terminologySimplifyRequest{
		ID: id,
	}
	parseJSON, err := json.Marshal(requestBody)
	if err != nil {
		return simpleTerms, errors.New("internal server error")
	}
	responseBody := bytes.NewBuffer(parseJSON)
	// Call Simplex service
	// TODO:  Endpoint can be fetched from configmap? localhost can be the service name
	resp, err := http.Post("http://"+SimplexServiceHost+":8083/v1/elaborate", "application/json", responseBody)
	if err != nil {
		return simpleTerms, errors.New("internal server error")
	}
	err = json.NewDecoder(resp.Body).Decode(&simpleTerms)
	if err != nil {
		return simpleTerms, errors.New("internal server error")
	}
	return simpleTerms, err
}
