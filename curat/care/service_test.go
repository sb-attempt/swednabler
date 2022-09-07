package curat

import (
	"context"
	"reflect"
	"testing"
)

func MockNewTerminologyService() *mockTerminologyService {
	return &mockTerminologyService{}
}

type mockTerminologyService struct{}

func (m mockTerminologyService) ValidateToken(token string) error {
	return nil
}

func TestNewTerminologyService(t *testing.T) {
	tests := []struct {
		name string
		want *terminologyService
	}{
		{
			name: "Test creation of terminology service",
			want: &terminologyService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTerminologyService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTerminologyService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_terminologyService_TerminologyList(t1 *testing.T) {
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    Terms
		wantErr bool
	}{
		{
			name: "Test terminology list service with invalid token",
			args: args{
				ctx:   context.Background(),
				token: "Bearer invalid_token",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := terminologyService{}
			got, err := t.TerminologyList(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t1.Errorf("TerminologyList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("TerminologyList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_terminologyService_TerminologySimplified(t1 *testing.T) {
	type args struct {
		ctx   context.Context
		token string
		id    int
	}
	tests := []struct {
		name    string
		args    args
		want    terminologySimplifyResponse
		wantErr bool
	}{
		{
			name: "Test terminology simplified with invalid token",
			args: args{
				ctx:   context.Background(),
				token: "Bearer invalid_token",
				id:    1,
			},
			want:    terminologySimplifyResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := terminologyService{}
			got, err := t.TerminologySimplified(tt.args.ctx, tt.args.token, tt.args.id)
			if (err != nil) != tt.wantErr {
				t1.Errorf("TerminologySimplified() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("TerminologySimplified() got = %v, want %v", got, tt.want)
			}
		})
	}
}
