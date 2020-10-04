package o47

import "testing"

func TestMaxGiftValue(t *testing.T) {
	type args struct {
		gifts [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{[][]int{
				{1, 10, 3, 8},
				{12, 2, 9, 6},
				{5, 7, 4, 11},
				{3, 7, 16, 5},
			}},
			want: 53,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxGiftValue(tt.args.gifts); got != tt.want {
				t.Errorf("MaxGiftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
