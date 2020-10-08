package o59

import (
	"reflect"
	"testing"
)

func TestMaxOfQueue(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{2, 3, 4, 2, 6, 2, 5, 1},
				k:    3,
			},
			want: []int{4, 4, 6, 6, 6, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxOfQueue(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxOfQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
