package leetcode

// 和层次遍历类似，利用栈实现
// 利用一个变量记录该层是否需要倒转
func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	queue := []*TreeNode{root}
	reverse := false
	for len(queue) > 0 {
		tmp := queue
		ret = append(ret, zigzagConvert(tmp, reverse))
		reverse = !reverse
		queue = []*TreeNode{}
		for _, node := range tmp {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return ret
}

func zigzagConvert(nodes []*TreeNode, reverse bool) []int {
	ret := []int{}
	if reverse {
		for i := len(nodes) - 1; i >= 0; i-- {
			ret = append(ret, nodes[i].Val)
		}
	} else {
		for i := 0; i < len(nodes); i++ {
			ret = append(ret, nodes[i].Val)
		}
	}
	return ret
}
