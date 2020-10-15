package leetcode

// recursive
func connect0(root *Node) *Node {
	if root == nil || root.Left == nil || root.Right == nil {
		return root
	}
	connect0(root.Left)
	connect0(root.Right)
	left := root.Left
	right := root.Right
	for left != nil && right != nil {
		left.Next = right
		left = left.Right
		right = right.Left
	}
	return root
}

//
func connect1(root *Node) *Node {
	if root == nil || root.Left == nil || root.Right == nil {
		return root
	}
	leftFirst := root
	for leftFirst != nil {
		tmp := leftFirst
		leftFirst = tmp.Left
		var prev *Node
		for ; tmp != nil; tmp = tmp.Next {
			if tmp.Left != nil {
				tmp.Left.Next = tmp.Right
			}
			if prev != nil {
				prev.Next = tmp.Left
			}
			prev = tmp.Right
		}
	}
	return root
}
