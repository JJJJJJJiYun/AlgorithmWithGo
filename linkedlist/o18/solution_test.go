package o18

import (
	"LeetCode/utils"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	type args struct {
		first  *utils.LinkedNode
		second *utils.LinkedNode
		third  *utils.LinkedNode
		fourth *utils.LinkedNode
		node   *utils.LinkedNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				first:  &utils.LinkedNode{Val: 1},
				second: &utils.LinkedNode{Val: 2},
				third:  &utils.LinkedNode{Val: 3},
				fourth: &utils.LinkedNode{Val: 4},
			},
		},
		{
			name: "case2",
			args: args{
				first:  &utils.LinkedNode{Val: 1},
				second: &utils.LinkedNode{Val: 2},
				third:  &utils.LinkedNode{Val: 3},
				fourth: &utils.LinkedNode{Val: 4},
			},
		},
		{
			name: "case3",
			args: args{
				first:  &utils.LinkedNode{Val: 1},
				second: &utils.LinkedNode{Val: 2},
				third:  &utils.LinkedNode{Val: 3},
				fourth: &utils.LinkedNode{Val: 4},
			},
		},
	}
	tests[0].args.node = tests[0].args.first
	tests[1].args.node = tests[1].args.second
	tests[2].args.node = tests[2].args.fourth
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.first.Next = tt.args.second
			tt.args.second.Next = tt.args.third
			tt.args.third.Next = tt.args.fourth
			got := DeleteNode(tt.args.first, tt.args.node)
			utils.PrintLinkedList(got)
		})
	}
}

func TestDeleteDuplicateNode(t *testing.T) {
	type args struct {
		first *utils.LinkedNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{first: &utils.LinkedNode{
				Val: 1,
				Next: &utils.LinkedNode{
					Val: 2,
					Next: &utils.LinkedNode{
						Val: 3,
						Next: &utils.LinkedNode{
							Val: 3,
							Next: &utils.LinkedNode{
								Val: 4,
							},
						},
					},
				},
			}},
		},
		{
			name: "case2",
			args: args{first: &utils.LinkedNode{
				Val: 1,
				Next: &utils.LinkedNode{
					Val: 2,
					Next: &utils.LinkedNode{
						Val: 3,
						Next: &utils.LinkedNode{
							Val: 4,
							Next: &utils.LinkedNode{
								Val: 4,
							},
						},
					},
				},
			}},
		},
		{
			name: "case3",
			args: args{first: &utils.LinkedNode{
				Val: 1,
				Next: &utils.LinkedNode{
					Val: 1,
					Next: &utils.LinkedNode{
						Val: 2,
						Next: &utils.LinkedNode{
							Val: 3,
							Next: &utils.LinkedNode{
								Val: 4,
							},
						},
					},
				},
			}},
		},
		{
			name: "case4",
			args: args{first: &utils.LinkedNode{
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
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeleteDuplicateNode(tt.args.first)
			utils.PrintLinkedList(got)
		})
	}
}
