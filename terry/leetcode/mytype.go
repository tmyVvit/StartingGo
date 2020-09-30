package leetcode

import "fmt"

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

// ListNode is a simple node for list/linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// BuildList is a func to build list by array
func BuildList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0], Next: nil}
	pre := head
	for i := 1; i < len(arr); i++ {
		node := &ListNode{Val: arr[i], Next: nil}
		pre.Next = node
		pre = node
	}
	return head
}

// PrintList is a func to print the list to stdout
func PrintList(head *ListNode) {
	if head == nil {
		return
	}
	for node := head; node != nil; node = node.Next {
		fmt.Printf("%d, ", node.Val)
	}
	fmt.Println()
}
