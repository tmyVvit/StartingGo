package leetcode

func inorderTraversal(root *TreeNode) (res []int) {
	var internal func(node *TreeNode)
	internal = func(node *TreeNode) {

		if root == nil {
			return
		}
		inorderTraversal(root.Left)
		res = append(res, root.Val)
		inorderTraversal(root.Right)
	}
	internal(root)
	return res
}
