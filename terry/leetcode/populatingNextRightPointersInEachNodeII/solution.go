package populatingNextRightPointersInEachNodeII

import (
	mytype "terry/leetcode/mytype"
)

// 使用队列解决该问题
func connect(root *mytype.Node) *mytype.Node {
	queue := []*mytype.Node{root}
	for len(queue) > 0 {
		size := len(queue)
		curr := queue
		queue = nil
		// 保证每次循环一层的节点
		for i := 0; i < size; i++ {
			node := curr[i]
			if i < size-1 {
				node.Next = curr[i+1]
			}
			if node == nil {
				continue
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}
