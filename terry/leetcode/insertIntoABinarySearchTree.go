package leetcode

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val, Left: nil, Right: nil}
	}
	node := root
	for {
		if node.Val > val {
			if node.Left == nil {
				node.Left = &TreeNode{Val: val, Left: nil, Right: nil}
				return root
			}
			node = node.Left
		} else {
			if node.Right == nil {
				node.Right = &TreeNode{Val: val, Left: nil, Right: nil}
				return root
			}
			node = node.Right
		}
	}
}
