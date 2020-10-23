package leetcode

// GetAllElements 合并两个二叉搜索树并排序
// 首先将两个tree转化成两个数组
// 然后数组合并
func GetAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1 := treeToSlice(root1)
	arr2 := treeToSlice(root2)
	return mergeSorted(arr1, arr2)
}

func treeToSlice(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := treeToSlice(root.Left)
	left = append(left, root.Val)
	return append(left, treeToSlice(root.Right)...)
}

func mergeSorted(arr1 []int, arr2 []int) (res []int) {
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		res = append(res, arr1[i])
		i++
	}
	for j < len(arr2) {
		res = append(res, arr2[j])
		j++
	}
	return res
}
