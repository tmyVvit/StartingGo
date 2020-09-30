package leetcode

import "fmt"

// func main() {
// 	root := TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{6, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{7, nil, nil}}}
// 	root.print()
// 	fmt.Println()
// 	convertBST(&root)
// 	root.print()
// }

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var middle func(root *TreeNode) *TreeNode
	middle = func(root *TreeNode) *TreeNode {
		if root != nil {
			// 使用反序中序遍历，从右往左计算
			middle(root.Right)
			sum += root.Val
			root.Val = sum
			middle(root.Left)
		}
		return root
	}
	return middle(root)
}

func (root *TreeNode) print() {
	fmt.Printf("%d ", root.Val)
	if root.Left != nil {
		root.Left.print()
	}
	if root.Right != nil {
		root.Right.print()
	}
}
