package leetcode

func lowestCommonAncestor(root, p, q *TreeNode) (res *TreeNode) {

	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil || node.Val == p.Val || node.Val == q.Val {
			return node
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left == nil {
			return right
		}
		if right == nil {
			return left
		}
		return node
	}
	return dfs(root)
}
