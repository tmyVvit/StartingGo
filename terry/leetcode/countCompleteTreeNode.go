package leetcode

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	height := 0
	tmp := root
	for tmp.Left != nil {
		height++
		tmp = tmp.Left
	}
	start, end := 1<<height, (1<<(height+1))-1
	for start < end {
		mid := (end-start)/2 + start
		if leafNodeExist(root, height, mid) {
			start = mid
		} else {
			end = mid - 1
		}
	}
	return start
}

func leafNodeExist(root *TreeNode, height int, leafBit int) bool {
	mask := 1 << (height - 1)
	node := root
	for ; mask > 0 && node != nil; mask >>= 1 {
		if leafBit&mask == 0 {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return node != nil
}
