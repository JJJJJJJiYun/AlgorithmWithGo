package bilibili

import "testing"

func TestFind(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				mat: [][]int{
					{0, 3, 4, 4, 5, 6, 9},
					{2, 4, 5, 9, 9, 10},
					{3, 5, 8, 9, 11},
					{1, 4, 5, 7, 9},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.mat); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
