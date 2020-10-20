package leetcode

func reorderList(head *ListNode) {
	arr := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		arr = append(arr, node)
	}
	length := len(arr)
	if length < 2 {
		return
	}
	i, j := 0, length-1
	pre := &ListNode{Val: 0, Next: nil}
	for i <= j {
		pre.Next = arr[i]
		pre.Next.Next = arr[j]
		pre = arr[j]
		i++
		j--
	}
	pre.Next = nil
}

func reorderList2(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// 1. 分为左右两个链表
	left, right := splitList(head)
	// 2. 将右链表翻转
	right = reverseList2(right)
	// 3. 合并链表
	mergeList(left, right)
}

func splitList(head *ListNode) (*ListNode, *ListNode) {
	prev, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	return head, slow
}

func mergeList(left, right *ListNode) *ListNode {
	list := &ListNode{Val: 0, Next: nil}
	tlist, tleft, tright := list, left, right
	for tleft != nil && tright != nil {
		ltmp, rtmp := tleft.Next, tright.Next
		tlist.Next = tleft
		tlist.Next.Next = tright
		tlist = tright
		tleft, tright = ltmp, rtmp
	}
	return list.Next
}
