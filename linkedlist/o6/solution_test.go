package o6

import "testing"

func TestReversePrintLinkedList(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{node: &Node{
				val: 1,
				next: &Node{
					val: 2,
					next: &Node{
						val:  3,
						next: nil,
					},
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReversePrintLinkedList(tt.args.node)
		})
	}
}
