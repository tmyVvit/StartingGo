package leetcode

// 可以考虑到另一题 21: Merge two sorted list
// 使用二分法和归并排序
func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 使用快慢指针将list切分成两个
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tmpRight := slow.Next
	slow.Next = nil
	sortLeft := sortList2(head)
	sortRight := sortList2(tmpRight)
	root := &ListNode{Val: 0, Next: nil}
	pre := root
	left, right := sortLeft, sortRight
	// 将sortLeft, sortRight两个顺序链表合并
	for left != nil && right != nil {
		if left.Val > right.Val {
			pre.Next = right
			right = right.Next
		} else {
			pre.Next = left
			left = left.Next
		}
		pre = pre.Next
	}
	if left != nil {
		pre.Next = left
	} else {
		pre.Next = right
	}

	return root.Next
}
