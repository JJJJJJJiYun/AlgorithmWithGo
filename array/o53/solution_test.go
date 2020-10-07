package o53

import (
	"testing"
)

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

func TestMissingNum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{nums: []int{0, 1, 3, 4, 5}},
			want: 2,
		},
		{
			name: "case2",
			args: args{nums: []int{1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MissingNum(tt.args.nums); got != tt.want {
				t.Errorf("MissingNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindNumMatchIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{nums: []int{-3, 1, 5, 6, 7}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindNumMatchIndex(tt.args.nums); got != tt.want {
				t.Errorf("FindNumMatchIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
