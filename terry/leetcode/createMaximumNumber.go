package leetcode

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	m, n := len(nums1), len(nums2)

	// 从第一个数组中最少可以选start个，最多可以选end个
	start, end := max(0, k-n), min(m, k)
	maxArray := make([]int, k)
	for i := start; i <= end; i++ {
		select1 := selectFromArray(nums1, i)
		select2 := selectFromArray(nums2, k-i)
		merged := mergeMaxArray(select1, select2)
		if arrayCompare(merged, 0, maxArray, 0) {
			maxArray = merged
		}
	}
	return maxArray
}

// 从nums数组中选择k个数，使组成的数字最大
func selectFromArray(nums []int, k int) []int {
	// 需要去除的个数
	delcount := len(nums) - k
	// ret 是一个单调栈，栈底到栈顶单调递减（从左到右）
	ret := make([]int, k)
	top := -1

	for _, num := range nums {
		// 当栈中有数，并且num大于栈顶元素并且还可以删除数时，删去栈顶元素
		for top >= 0 && ret[top] < num && delcount > 0 {
			top--
			delcount--
		}
		if top < k-1 {
			//当栈中还没满时，入栈
			top++
			ret[top] = num
		} else {
			// 如果栈满了
			delcount--
		}
	}
	return ret
}

func mergeMaxArray(nums1 []int, nums2 []int) []int {
	// 合并成最大的数组
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	merged := make([]int, len(nums1)+len(nums2))
	i, j, k := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if arrayCompare(nums1, i, nums2, j) {
			merged[k] = nums1[i]
			k++
			i++
		} else {
			merged[k] = nums2[j]
			k++
			j++
		}
	}
	for i < len(nums1) {
		merged[k] = nums1[i]
		k++
		i++
	}
	for j < len(nums2) {
		merged[k] = nums2[j]
		k++
		j++
	}
	return merged
}

func arrayCompare(nums1 []int, idx1 int, nums2 []int, idx2 int) bool {
	for idx1 < len(nums1) && idx2 < len(nums2) {
		if nums1[idx1] == nums2[idx2] {
			idx1++
			idx2++
		} else {
			return nums1[idx1] > nums2[idx2]
		}
	}
	// 当数组全部相同时，更长的数组更大
	return (len(nums1) - idx1) > (len(nums2) - idx2)
}
