package invertBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var copyLeft *TreeNode
	var copyRight *TreeNode
	copyLeft = root.Left
	copyRight = root.Right
	root.Left = copyRight
	root.Right = copyLeft
	invertTree(root.Left)
	invertTree(root.Right)
	return root

}
