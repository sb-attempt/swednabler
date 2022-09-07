package open

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewJwtService(t *testing.T) {
	tests := []struct {
		name string
		want *jwtService
	}{
		{
			name: "Test JWT service creation",
			want: &jwtService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJwtService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJwtService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwtService_GetToken(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test JWT service to get token",
			args: args{
				username: "user1",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := jwtService{}
			got, err := j.GetToken(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.want {
				t.Errorf("GetToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwtService_ValidateToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
	}{
		{
			name: "Check for bad token",
			args: args{
				token: "some_token",
			},
			want:  false,
			want1: http.StatusBadRequest,
		},
		{
			name: "Check for invalid token",
			args: args{
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyMSIsImV4cCI6MTY2MjM2MjY2MH0.IZx0hNBI00KQ46WzqpsoodpEj6quwp3AZJYMrFLvoO0",
			},
			want:  false,
			want1: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := jwtService{}
			got, got1 := j.ValidateToken(tt.args.token)
			if got != tt.want {
				t.Errorf("ValidateToken() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ValidateToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
