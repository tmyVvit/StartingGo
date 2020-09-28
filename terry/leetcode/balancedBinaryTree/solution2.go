package balancedBinaryTree

import (
	"terry/leetcode/mytype"
)

// 在计算高度的同时判断是否是平衡的
func isBalanced2(root *mytype.TreeNode) bool {
	return height2(root) > -1
}

func height2(root *mytype.TreeNode) int {
	if root == nil {
		return 0
	}

	left := height2(root.Left)
	if left == -1 {
		return -1
	}
	right := height2(root.Right)

	if right == -1 || abs(left-right) > 1 {
		return -1
	}
	return max(left, right) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
