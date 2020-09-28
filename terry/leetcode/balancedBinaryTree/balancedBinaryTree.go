package balancedBinaryTree

import (
	"terry/leetcode/mytype"
)

func isBalanced(root *mytype.TreeNode) bool {
	if root == nil {
		return true
	}

	left := isBalanced(root.Left)
	if !left {
		return false
	}
	right := isBalanced(root.Right)
	if right {
		if abs(height(root.Left)-height(root.Right)) > 1 {
			return false
		}
		return true
	}
	return false
}

func height(root *mytype.TreeNode) (res int) {
	if root == nil {
		return
	}
	queue := []*mytype.TreeNode{root}
	for len(queue) > 0 {
		tmp := queue
		queue = nil
		for _, node := range tmp {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res++
	}

	return res
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
