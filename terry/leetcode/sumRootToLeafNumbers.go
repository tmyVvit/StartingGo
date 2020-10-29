package leetcode

type myNode struct {
	node *TreeNode
	flag int
}

func SumNumbers(root *TreeNode) int {
	st := []*myNode{}
	sum := []int{}
	node := root
	for node != nil || len(st) > 0 {
		for node != nil {
			st = append(st, &myNode{node, 0})
			node = node.Left
		}
		node = st[len(st)-1].node
		if node.Right == nil {
			sum = append(sum, getTreeSum(st))
			topIdx := len(st) - 2
			for topIdx >= 0 && (st[topIdx].flag == 1 || st[topIdx].node.Right == nil) {
				topIdx--
			}
			if topIdx < 0 {
				node = nil
				st = st[:0]
			} else {
				st = st[:topIdx+1]
				node = st[topIdx].node.Right
				st[topIdx].flag = 1
			}

		} else {
			st[len(st)-1].flag = 1
			node = node.Right
		}
	}
	return getSum(sum)
}

func getTreeSum(nodes []*myNode) int {
	sum := 0
	for _, node := range nodes {
		sum = sum*10 + node.node.Val
	}
	return sum
}

func getSum(list []int) int {
	sum := 0
	for _, val := range list {
		sum += val
	}
	return sum
}

func SumNumbersI(root *TreeNode) int {
	var dfs func(node *TreeNode, presum int) int
	dfs = func(node *TreeNode, presum int) int {
		if node == nil {
			return 0
		}
		sum := presum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return sum
		}
		return dfs(node.Left, sum) + dfs(node.Right, sum)

	}
	return dfs(root, 0)
}
