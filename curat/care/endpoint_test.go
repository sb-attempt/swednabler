package curat

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"reflect"
	"testing"
)

// Mock JWT service through mock interface
type MockTerminologyService struct{}

func (m MockTerminologyService) ValidateToken(token string) error {
	return nil
}

func (m MockTerminologyService) TerminologyList(ctx context.Context, token string) (Terms, error) {
	term := Terms{}
	return term, nil
}

func (m MockTerminologyService) TerminologySimplified(ctx context.Context, token string, id int) (terminologySimplifyResponse, error) {
	simple := terminologySimplifyResponse{
		SimpleTerms: SimpleTerms{
			Name:        "1",
			FullForm:    "know your customer",
			Description: "some description",
		},
		Err: "",
	}
	return simple, nil
}

func NewMockTerminologyService() *MockTerminologyService {
	return &MockTerminologyService{}
}

func Test_terminologyListEndpoint(t *testing.T) {
	type args struct {
		svc TerminologyService
	}
	tests := []struct {
		name string
		args args
		want endpoint.Endpoint
	}{
		{
			name: "Test terminology list endpoint",
			args: args{
				svc: NewMockTerminologyService(),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := terminologyListEndpoint(tt.args.svc); reflect.DeepEqual(got, tt.want) {
				t.Errorf("terminologyListEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_terminologySimplifyEndpoint(t *testing.T) {
	type args struct {
		svc TerminologyService
	}
	tests := []struct {
		name string
		args args
		want endpoint.Endpoint
	}{
		{
			name: "Test Simplify endpoint",
			args: args{
				svc: NewMockTerminologyService(),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := terminologySimplifyEndpoint(tt.args.svc); reflect.DeepEqual(got, tt.want) {
				t.Errorf("terminologySimplifyEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
