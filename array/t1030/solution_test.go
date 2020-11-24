package t1030

import (
	"reflect"
	"testing"
)

func Test_allCellsDistOrder2(t *testing.T) {
	type args struct {
		n  int
		m  int
		r0 int
		c0 int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				n:  1,
				m:  2,
				r0: 0,
				c0: 0,
			},
			want: [][]int{{0, 0}, {0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allCellsDistOrder2(tt.args.n, tt.args.m, tt.args.r0, tt.args.c0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allCellsDistOrder2() = %v, want %v", got, tt.want)
			}
		})
	}
}
