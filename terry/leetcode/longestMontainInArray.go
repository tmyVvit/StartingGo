package leetcode

// LongestMountain 返回最长山脉子数组长度
func LongestMountain(A []int) int {
	if len(A) <= 2 {
		return 0
	}
	top := []int{0}
	// 首先找到所有山脉的顶点的下标，存入top数组
	for i := 1; i < len(A)-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			top = append(top, i)
		}
	}
	top = append(top, len(A)-1)
	maxLength := 0
	// 从山脉顶点开始分别向左向右计算山脉长度
	// 山脉top[i]的长度必然小于top[i+1] - top[i-1]
	for i := 1; i < len(top)-1; i++ {
		tmpMax := 1
		for j := top[i]; j > top[i-1]; j-- {
			if A[j] > A[j-1] {
				tmpMax++
			} else {
				break
			}
		}
		for j := top[i]; j < top[i+1]; j++ {
			if A[j] > A[j+1] {
				tmpMax++
			} else {
				break
			}
		}
		maxLength = max(maxLength, tmpMax)
	}
	return maxLength
}

// LongestMountain2 使用动态规划计算最长山脉的长度
func LongestMountain2(A []int) int {
	if len(A) < 3 {
		return 0
	}
	// left[i], right[i] 分别代表以i为山顶到左山脚/右山脚的距离
	// 如果 A[i] > A[i-1]，那么left[i] = left[i-1] + 1，否则left[i] = 0
	// 同理 A[i] > A[i+1]，那么right[i] = right[i+1] + 1，否则right[i] = 0
	// 只有当left[i]>0且right[i]>0时，i才可能是山顶
	left, right := make([]int, len(A)), make([]int, len(A))

	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 0
		}
	}
	for i := len(A) - 2; i >= 0; i-- {
		if A[i] > A[i+1] {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 0
		}
	}

	maxLength := 0
	for i := 0; i < len(A); i++ {
		if left[i] == 0 || right[i] == 0 {
			continue
		}
		maxLength = max(maxLength, left[i]+right[i]+1)
	}
	return maxLength
}
