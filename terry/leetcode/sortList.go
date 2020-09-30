package leetcode

import (
	"math"
)

// func main() {
// 	head := mytype.BuildList([]int{-1, 5, 3, 4, 0})
// 	mytype.PrintList(sortList(head))
// }

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	root := &ListNode{Val: math.MinInt32, Next: head}
	pre, curr := head, head.Next
	for curr != nil {
		if curr.Val > pre.Val {
			pre = curr
			curr = curr.Next
			continue
		}
		for node := root; node != curr; node = node.Next {
			if node.Next.Val < curr.Val {
				continue
			}
			pre.Next = curr.Next
			curr.Next = node.Next
			node.Next = curr
			curr = pre.Next
			break
		}
	}
	return root.Next
}
