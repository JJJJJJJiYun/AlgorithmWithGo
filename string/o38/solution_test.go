package o38

import "testing"

func TestStringPermutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{s: "abc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StringPermutation(tt.args.s)
		})
	}
}

func TestStringPermutationWithRepeated(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{s: "abbc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StringPermutationWithRepeated(tt.args.s)
		})
	}
}
