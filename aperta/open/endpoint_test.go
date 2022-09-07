package open

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

// Mock JWT service through mock interface
type MockJwtService struct{}

func NewMockJwtService() *MockJwtService {
	return &MockJwtService{}
}

// Implement methods for the mock interface that is implemented by the mock struct.
func (m MockJwtService) GetToken(username string) (string, error) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyMSIsImV4cCI6MTY2MjM2MjY2MH0.IZx0hNBI00KQ46WzqpsoodpEj6quwp3AZJYMrFLvoO0"
	return token, nil
}

func (m MockJwtService) ValidateToken(token string) (bool, int) {
	tokenExpected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyMSIsImV4cCI6MTY2MjM2MjY2MH0.IZx0hNBI00KQ46WzqpsoodpEj6quwp3AZJYMrFLvoO0"
	if token != tokenExpected {
		return false, http.StatusUnauthorized
	}
	return true, http.StatusOK
}

func Test_authenticateEndpoint(t *testing.T) {
	type args struct {
		svc Service
		jwt JwtService
		req authenticateRequest
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test authentication Endpoint with valid user",
			args: args{
				svc: NewService(),
				jwt: NewMockJwtService(),
				req: authenticateRequest{
					Username: "user1",
					Password: "password1",
				},
			},
			want: "",
		},

		{
			name: "Test authentication Endpoint with invalid user",
			args: args{
				svc: NewService(),
				jwt: NewMockJwtService(),
				req: authenticateRequest{
					Username: "user10",
					Password: "password10",
				},
			},
			want: "the user is invalid",
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			e := authenticateEndpoint(tt.args.svc, tt.args.jwt)
			authResponse, _ := e(context.Background(), tt.args.req)
			got := authResponse.(authenticateResponse).Err
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authorizeEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authorizeEndpoint(t *testing.T) {
	type args struct {
		svc Service
		jwt JwtService
		req authorizeRequest
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test authorization Endpoint with valid token",
			args: args{
				svc: NewService(),
				jwt: NewMockJwtService(),
				req: authorizeRequest{
					Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyMSIsImV4cCI6MTY2MjM2MjY2MH0.IZx0hNBI00KQ46WzqpsoodpEj6quwp3AZJYMrFLvoO0",
				},
			},
			want: http.StatusOK,
		},

		{
			name: "Test authorization Endpoint with invalid token",
			args: args{
				svc: NewService(),
				jwt: NewMockJwtService(),
				req: authorizeRequest{
					Token: "invalid",
				},
			},
			want: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			e := authorizeEndpoint(tt.args.svc, tt.args.jwt)
			authResponse, _ := e(context.Background(), tt.args.req)
			got := authResponse.(authorizeResponse).ErrCode
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authorizeEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
