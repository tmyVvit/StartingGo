package leetcode

import (
	"math"
)

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	fastSort(nums)
	maxGap := math.MinInt32
	for i := 1; i < len(nums); i++ {
		tmp := nums[i] - nums[i-1]
		if maxGap < tmp {
			maxGap = tmp
		}
	}
	return maxGap
}

// 快速排序
func fastSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	start, end := 0, len(nums)-1
	pivot := nums[start]
	for start < end {
		for ; nums[end] >= pivot && start < end; end-- {
		}
		nums[start] = nums[end]
		for ; nums[start] <= pivot && start < end; start++ {
		}
		nums[end] = nums[start]
	}
	nums[start] = pivot
	fastSort(nums[:start])
	fastSort(nums[start+1:])
}

// 基数排序 LSD 最低有效位优先
func lsd(nums []int) {
	maxnum := math.MinInt32
	for _, num := range nums {
		maxnum = max(maxnum, num)
	}
	copynum := make([]int, len(nums))
	for exp := 1; exp <= maxnum; exp *= 10 {
		index := make([]int, 10)
		// 计算排序后的下标
		for _, num := range nums {
			index[(num/exp)%10]++
		}
		for i := 1; i < 10; i++ {
			index[i] += index[i-1]
		}
		// 因为使用了index来记录下标，相同基数的数中我们先填后面的数，所以需要从后往前
		for i := len(nums) - 1; i >= 0; i-- {
			tmp := (nums[i] / exp) % 10
			//相同基数的数中我们先填后面的数
			copynum[index[tmp]-1] = nums[i]
			index[tmp]--
		}
		copy(nums, copynum)
	}

}
