package leetcode

// 两次二分查找
func searchRange(nums []int, target int) (ret []int) {
	length := len(nums)
	if length == 0 || target < nums[0] || target > nums[length-1] {
		return []int{-1, -1}
	}
	start, end := 0, length-1
	for start < end {
		mid := (start + end) / 2
		if nums[mid] == target {
			end = mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	if nums[start] == target {
		ret = append(ret, start)
	} else {
		return []int{-1, -1}
	}
	// 第二次可以缩小范围
	start, end = ret[0], length-1
	for start < end {
		mid := (start + end + 1) / 2
		if nums[mid] == target {
			start = mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return append(ret, end)
}
