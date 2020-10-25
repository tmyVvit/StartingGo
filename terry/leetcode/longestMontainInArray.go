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
