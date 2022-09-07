package simplex

import (
	"github.com/go-kit/kit/endpoint"
	"reflect"
	"testing"
)

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
			name: "Test terminology endpoint",
			args: args{
				svc: NewTerminologyService(),
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
