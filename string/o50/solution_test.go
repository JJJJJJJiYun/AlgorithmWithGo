package o50

import "testing"

func TestFirstAppearOnceChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{s: "abaccdeff"},
			want: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstAppearOnceChar(tt.args.s); got != tt.want {
				t.Errorf("FirstAppearOnceChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
