package simplex

import (
	"context"
	"reflect"
	"testing"
)

func Test_terminologyService_TerminologySimplified(t1 *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		want    SimpleTerms
		wantErr bool
	}{
		{
			name: "Test terminology simplified service",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: SimpleTerms{
				Name:        "",
				FullForm:    "",
				Description: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := terminologyService{}
			got, err := t.TerminologySimplified(tt.args.ctx, tt.args.id)
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
