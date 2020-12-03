package leetcode

// 使用二分查找，如果左子树节点个数大于k-1个，就去左子树找
// 如果等于 k-1 个就返回根节点
// 如果小于 k-1 个就去右子树 找右子树中第 k-1-L 个(从1开始)节点 （左子树节点数L）
func kthSmallest(root *TreeNode, k int) int {
	// 使用map来记录以该节点为根节点的数的节点个数，防止重复计算
	nodesCount := make(map[int]int)
	var treeNodesCount func(node *TreeNode) int
	treeNodesCount = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if count, ok := nodesCount[node.Val]; ok {
			return count
		}
		left := treeNodesCount(node.Left)
		right := treeNodesCount(node.Right)
		nodesCount[node.Val] = left + right + 1
		return nodesCount[node.Val]
	}
	var kSmallest func(node *TreeNode, t int) int
	kSmallest = func(node *TreeNode, t int) int {
		leftCount := treeNodesCount(node.Left)
		if leftCount == t-1 {
			return node.Val
		} else if leftCount < t-1 {
			return kSmallest(node.Right, t-1-leftCount)
		} else {
			return kSmallest(node.Left, t)
		}
	}
	return kSmallest(root, k)
}
