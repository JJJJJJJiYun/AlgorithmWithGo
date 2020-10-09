package o25

import (
	"AlgorithWithGo/utils"
	"testing"
)

func TestMergeTwoSortedLinkedList(t *testing.T) {
	type args struct {
		first1 *utils.LinkedNode
		first2 *utils.LinkedNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				first1: &utils.LinkedNode{
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
				first2: &utils.LinkedNode{
					Val: 6,
					Next: &utils.LinkedNode{
						Val: 7,
						Next: &utils.LinkedNode{
							Val: 8,
							Next: &utils.LinkedNode{
								Val: 9,
								Next: &utils.LinkedNode{
									Val: 10,
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
			utils.PrintLinkedList(MergeTwoSortedLinkedList(tt.args.first1, tt.args.first2))
		})
	}
}
