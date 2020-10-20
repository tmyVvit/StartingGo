package leetcode

func reverseList(head *ListNode) *ListNode {
	var reverse func(node *ListNode) (*ListNode, *ListNode)
	reverse = func(node *ListNode) (*ListNode, *ListNode) {
		if node == nil || node.Next == nil {
			return node, node
		}
		nHead, tail := reverse(node.Next)
		tail.Next = node
		tail = tail.Next
		node.Next = nil
		return nHead, tail
	}
	list, _ := reverse(head)
	return list
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	next := head
	for next != nil {
		tmp := next.Next
		next.Next = prev
		prev = next
		next = tmp
	}
	return prev
}
