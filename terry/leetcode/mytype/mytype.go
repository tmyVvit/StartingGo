package mytype

// TreeNode is a simple tree node struct.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Node is a simple tree node struct with next pointer.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
