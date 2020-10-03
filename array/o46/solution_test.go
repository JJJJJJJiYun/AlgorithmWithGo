package o46

import "testing"

func TestWayOfTransNums(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{num: 12258},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WayOfTransNums(tt.args.num); got != tt.want {
				t.Errorf("WayOfTransNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
