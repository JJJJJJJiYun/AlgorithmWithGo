package t942

import (
	"reflect"
	"testing"
)

func Test_diStringMatch(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{S: "DIIDD"},
			want: []int{5, 0, 1, 4, 3, 2},
		},
		{
			name: "case2",
			args: args{S: "III"},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "case3",
			args: args{S: "DDD"},
			want: []int{3, 2, 1, 0},
		},
		{
			name: "case4",
			args: args{S: "IDID"},
			want: []int{0, 4, 1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diStringMatch(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
