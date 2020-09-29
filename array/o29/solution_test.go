package o29

import (
	"testing"
)

func TestPrintMatrix(t *testing.T) {
	type args struct {
		m [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{m: [][]int{
				{1, 2, 3, 4},
				{10, 11, 12, 5},
				{9, 8, 7, 6},
			}},
		},
		{
			name: "case2",
			args: args{m: [][]int{
				{1, 2, 3},
				{10, 11, 4},
				{9, 12, 5},
				{8, 7, 6},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintMatrix(tt.args.m)
		})
	}
}
