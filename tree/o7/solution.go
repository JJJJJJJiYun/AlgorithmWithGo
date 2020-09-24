package o7

import (
	"LeetCode/utils"
	"fmt"
)

// 重建二叉树，前序+中序

// 1,2,4,7,3,5,6,8 pre
// 4,7,2,1,5,3,8,6 in
func ReconstructTreeWithPreAndInOrderTraversal(pre, in []int) *utils.TreeNode {
	var helper func(pre, in []int, ps, pe, is, ie int) *utils.TreeNode
	helper = func(pre, in []int, ps, pe, is, ie int) *utils.TreeNode {
		if ps > pe {
			return nil
		}
		node := &utils.TreeNode{Val: pre[ps]}
		nodeIndex := -1
		for i := is; i <= ie; i++ {
			if in[i] == node.Val {
				nodeIndex = i
				break
			}
		}
		if nodeIndex == -1 {
			panic(fmt.Errorf("invalid input"))
		}
		node.Left = helper(pre, in, ps+1, ps+nodeIndex-is, is, nodeIndex-1)
		node.Right = helper(pre, in, ps+nodeIndex-is+1, pe, nodeIndex+1, ie)
		return node
	}
	return helper(pre, in, 0, len(pre)-1, 0, len(in)-1)
}
