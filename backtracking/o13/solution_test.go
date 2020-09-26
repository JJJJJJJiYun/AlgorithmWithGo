package o13

import "testing"

func Test_getNumSum(t *testing.T) {
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
			args: args{num: 12},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumSum(tt.args.num); got != tt.want {
				t.Errorf("getNumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRobotMove(t *testing.T) {
	type args struct {
		rows      int
		cols      int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				rows:      3,
				cols:      3,
				threshold: 2,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RobotMove(tt.args.rows, tt.args.cols, tt.args.threshold); got != tt.want {
				t.Errorf("RobotMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
