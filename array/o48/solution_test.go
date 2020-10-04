package o48

import "testing"

func TestLengthOfMaxSubStringWithNoRepeated(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{s: "arabcacfr"},
			want: 4,
		},
		{
			name: "case2",
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "case3",
			args: args{s: "bbbbbb"},
			want: 1,
		},
		{
			name: "case4",
			args: args{s: "pwwkew"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfMaxSubStringWithNoRepeated(tt.args.s); got != tt.want {
				t.Errorf("LengthOfMaxSubStringWithNoRepeated() = %v, want %v", got, tt.want)
			}
		})
	}
}
