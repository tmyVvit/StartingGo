package leetcode

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	// 首先根据 区间左边界由小到大（如果相等，则右边界从大到小）排序
	// 确保只能是前面的区间覆盖后面的区间
	sort.Slice(intervals, func(a, b int) bool {
		if intervals[a][0] == intervals[b][0] {
			return intervals[a][1] > intervals[b][1]
		}
		return intervals[a][0] < intervals[b][0]
	})
	pre, count := 0, 0
	for i := 1; i < len(intervals); i++ {
		if isCovered(intervals[i], intervals[pre]) {
			count++
		} else {
			pre = i
		}
	}
	return len(intervals) - count
}

func isCovered(arr1 []int, arr2 []int) bool {
	return arr1[0] >= arr2[0] && arr1[1] <= arr2[1]
}
