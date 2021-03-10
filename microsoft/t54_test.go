package microsoft

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			name: "case1",
			args: args{matrix: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {11, 12, 13, 14, 15, 16, 17, 18, 19, 20}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := spiralOrder(tt.args.matrix); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("spiralOrder() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
