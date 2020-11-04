package leetcode

import "sort"

func merge(intervals [][]int) (ret [][]int) {
	if len(intervals) == 0 {
		return intervals
	}
	// 将所有区间根据每个区间的开始排序，升序
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	end := intervals[0][1]
	start := intervals[0][0]
	//  [[1,3],[2,6],[8,10],[15,18]]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][1] > end && intervals[i][0] <= end {
			// 当前的区间和之前的区间有交集,区间可以合并
			end = intervals[i][1]
		} else if intervals[i][1] <= end {
			// 之前区间已经包含了当前区间
			continue
		} else if intervals[i][0] > end {
			// 区间之间已经断开
			ret = append(ret, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}
	return ret
}

func mergeI(intervals [][]int) (ret [][]int) {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	for i := 0; i < len(intervals); {
		start, end := intervals[i][0], intervals[i][1]
		j := i + 1
		// 因为已将按照区间开始排序了，所以不需要再判断 intervals[j][1] 与 end的大小
		// 只要intervals[j][0] <= end 那么就肯定有交集，可以合并
		for j < len(intervals) && intervals[j][0] <= end {
			end = max(end, intervals[j][1])
			j++
		}
		ret = append(ret, []int{start, end})
		i = j
	}
	return ret
}
