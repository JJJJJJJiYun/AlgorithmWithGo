package microsoft

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{
			name: "case1",
			args: args{n: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := generateParenthesis(tt.args.n); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("generateParenthesis() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
