package o53

import (
	"fmt"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	fmt.Println(missingNumber([]int{0, 1, 3, 4, 5}))
}

func Test_findLeftBorder(t *testing.T) {
	type args struct {
		nums []int
		num  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 1, 2, 3, 3, 3, 4, 5, 5},
				num:  3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLeftBorder(tt.args.nums, tt.args.num); got != tt.want {
				t.Errorf("findLeftBorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRightBorder(t *testing.T) {
	type args struct {
		nums []int
		num  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 1, 2, 3, 3, 3, 4, 5, 5},
				num:  3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRightBorder(tt.args.nums, tt.args.num); got != tt.want {
				t.Errorf("findRightBorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountNumInSortedArray(t *testing.T) {
	type args struct {
		nums []int
		num  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 1, 2, 3, 3, 3, 4, 5, 5},
				num:  3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountNumInSortedArray(tt.args.nums, tt.args.num); got != tt.want {
				t.Errorf("CountNumInSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
