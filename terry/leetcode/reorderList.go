package leetcode

func ReorderList(head *ListNode) {
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
