package o31

import "testing"

func TestIsPopOfPush(t *testing.T) {
	type args struct {
		push []int
		pop  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				push: []int{1, 2, 3, 4, 5},
				pop:  []int{3, 2, 4, 5, 1},
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				push: []int{1, 2, 3, 4, 5},
				pop:  []int{3, 2, 5, 4, 1},
			},
			want: true,
		},
		{
			name: "case3",
			args: args{
				push: []int{1, 2, 3, 4, 5},
				pop:  []int{1, 2, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				push: []int{1, 2, 3, 4, 5},
				pop:  []int{4, 3, 5, 1, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPopOfPush(tt.args.push, tt.args.pop); got != tt.want {
				t.Errorf("IsPopOfPush() = %v, want %v", got, tt.want)
			}
		})
	}
}
