package open

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	tests := []struct {
		name string
		want *service
	}{
		{
			name: "Test new service creation",
			want: &service{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Authenticate(t *testing.T) {
	type args struct {
		ctx      context.Context
		jwt      JwtService
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test authentication service valid user",
			args: args{
				ctx:      context.Background(),
				jwt:      MockJwtService{},
				username: "user1",
				password: "password1",
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyMSIsImV4cCI6MTY2MjM2MjY2MH0.IZx0hNBI00KQ46WzqpsoodpEj6quwp3AZJYMrFLvoO0",
			wantErr: false,
		},
		{
			name: "Test authentication service invalid user",
			args: args{
				ctx:      context.Background(),
				jwt:      MockJwtService{},
				username: "user10",
				password: "password10",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{}
			got, err := s.Authenticate(tt.args.ctx, tt.args.jwt, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Authenticate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Authorize(t *testing.T) {
	type args struct {
		ctx   context.Context
		jwt   JwtService
		token string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
	}{
		{
			name: "Test authorize service : valid token",
			args: args{
				ctx:   context.Background(),
				jwt:   MockJwtService{},
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyMSIsImV4cCI6MTY2MjM2MjY2MH0.IZx0hNBI00KQ46WzqpsoodpEj6quwp3AZJYMrFLvoO0",
			},
			want:  true,
			want1: http.StatusOK,
		},
		{
			name: "Test authorize service : invalid token",
			args: args{
				ctx:   context.Background(),
				jwt:   MockJwtService{},
				token: "some_token",
			},
			want:  false,
			want1: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{}
			got, got1 := s.Authorize(tt.args.ctx, tt.args.jwt, tt.args.token)
			if got != tt.want {
				t.Errorf("Authorize() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Authorize() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
