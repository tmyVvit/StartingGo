package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, &TreeNode{4, nil, nil}}}
	fmt.Println(findMode(tree))
}

func findMode(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	var find func(node *TreeNode)
	var maxCounts, curr, preVal int
	flag := 0
	find = func(node *TreeNode) {
		if node == nil {
			return
		}
		find(node.Left)
		if flag == 0 {
			preVal = node.Val
			curr = 1
			flag++
		} else {
			if preVal == node.Val {
				curr++
			} else {
				if curr == maxCounts {
					res = append(res, preVal)
					curr = 1
				} else if curr < maxCounts {
					curr = 1
				} else {
					res = res[:0]
					res = append(res, preVal)
					maxCounts = curr
					curr = 1
				}
			}
			preVal = node.Val
		}
		find(node.Right)
	}

	find(root)
	if curr == maxCounts {
		res = append(res, preVal)
	} else if curr > maxCounts {
		res = []int{preVal}
	}
	return res
}
