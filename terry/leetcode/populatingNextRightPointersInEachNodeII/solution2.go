package populatingNextRightPointersInEachNodeII

import (
	mytype "terry/leetcode/mytype"
)

// 使用队列解决该问题
// 可以考虑当第i层已经有next时，就可以使用next遍历该层节点
// 在第x层时，构建第x+1层的next
func connect2(root *mytype.Node) *mytype.Node {
	first := root
	for first != nil {
		var pre *mytype.Node
		tmp := first
		first = nil
		for next := tmp; next != nil; next = next.Next {
			if next.Left == nil && next.Right == nil {
				continue
			}
			if next.Left != nil {
				if first == nil {
					first = next.Left
				}
				if pre != nil {
					pre.Next = next.Left
				}
				if next.Right != nil {
					next.Left.Next = next.Right
					pre = next.Right
				} else {
					pre = next.Left
				}
			} else {
				if first == nil {
					first = next.Right
				}
				if pre != nil {
					pre.Next = next.Right
				}
				pre = next.Right
			}
		}
	}

	return root
}
