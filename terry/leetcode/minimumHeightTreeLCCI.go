package leetcode

// Given a sorted (increasing order) array with unique integer elements, write an algo­rithm to create a binary search tree with minimal height.
func sortedArrayToBST(nums []int) *TreeNode {
	var createTree func(arr []int) *TreeNode
	createTree = func(arr []int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		if len(arr) == 1 {
			return &TreeNode{Val: arr[0], Left: nil, Right: nil}
		}
		mid := len(arr) / 2
		return &TreeNode{Val: arr[mid], Left: createTree(arr[:mid]), Right: createTree(arr[mid+1:])}
	}
	return createTree(nums)
}

func sortedArrayToBSTI(nums []int) *TreeNode {
	size := len(nums)
	if size == 0 {
		return nil
	}
	mid := size / 2
	return &TreeNode{Val: nums[mid], Left: sortedArrayToBSTI(nums[:mid]), Right: sortedArrayToBSTI(nums[mid+1:])}
}
