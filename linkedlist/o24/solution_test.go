package o24

import (
	"AlgorithWithGo/utils"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	type args struct {
		first *utils.LinkedNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				first: &utils.LinkedNode{
					Val: 1,
					Next: &utils.LinkedNode{
						Val: 2,
						Next: &utils.LinkedNode{
							Val: 3,
							Next: &utils.LinkedNode{
								Val: 4,
								Next: &utils.LinkedNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseLinkedList(tt.args.first)
			utils.PrintLinkedList(got)
		})
	}
}
