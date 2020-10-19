package t1576

import (
	"fmt"
	"testing"
)

func Test_modifyString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{s: "j?qg??b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := modifyString(tt.args.s)
			fmt.Println(got)
		})
	}
}
