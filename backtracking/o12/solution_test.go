package o12

import "testing"

func TestPathInMatrix(t *testing.T) {
	type args struct {
		matrix [][]rune
		target []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				matrix: [][]rune{
					{'a', 'b', 't', 'g'},
					{'c', 'f', 'c', 's'},
					{'j', 'd', 'e', 'h'},
				},
				target: []rune{'b', 'f', 'c', 'e'},
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				matrix: [][]rune{
					{'a', 'b', 't', 'g'},
					{'c', 'f', 'c', 's'},
					{'j', 'd', 'e', 'h'},
				},
				target: []rune{'a', 'b', 'f', 'b'},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PathInMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("PathInMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
