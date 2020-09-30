package leetcode

// 中序   左->中->右
// 前序   中->左->右
// 后序   左->右->中

// func main() {
// 	inorder := []int{9, 3, 15, 20, 7}
// 	postorder := []int{9, 15, 7, 20, 3}
// 	buildTree(inorder, postorder)
// }

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	counts := len(postorder)
	// 后序遍历的最后位就是根节点
	rootVal := postorder[counts-1]
	// 根据根节点将 中/后序遍历分为左右子树的中/后序遍历
	leftIn, leftPost, rightIn, rightPost := splitChildren(inorder, postorder, rootVal)
	left := buildTree(leftIn, leftPost)
	right := buildTree(rightIn, rightPost)
	root := new(TreeNode)
	root.Val = rootVal
	root.Left = left
	root.Right = right
	return root
}

func splitChildren(inorder []int, postorder []int, rootVal int) (leftIn, leftPost, rightIn, rightPost []int) {
	var leftCount int
	for index, val := range inorder {
		if val == rootVal {
			leftCount = index
			break
		}
	}
	leftIn = inorder[:leftCount]
	rightIn = inorder[leftCount+1:]
	leftPost = postorder[:leftCount]
	rightPost = postorder[leftCount : len(postorder)-1]
	return
}
