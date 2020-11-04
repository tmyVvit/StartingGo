package leetcode

func insertIntervals(intervals [][]int, newInterval []int) (ret [][]int) {
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}
	start, end := newInterval[0], newInterval[1]
	for i := 0; i < len(intervals); i++ {
		if start <= intervals[i][1] { //  注意不能用start <= intervals[i][0],因为当start > intervals[i][0]时可能start <= intervals[i][1]，此时也是有交集的
			// 此时找到第一个可能有交集的区间
			// 首先设置start
			start = min(start, intervals[i][0])
			// 由于还没有判断当前区间的结束与end的大小，所以设置j=i
			j := i
			for j < len(intervals) && intervals[j][0] <= end {
				end = max(end, intervals[j][1])
				j++
			}
			// 有交集的区间已经合并，加入结果集
			ret = append(ret, []int{start, end})
			// 将剩余的区间加入结果集
			for ; j < len(intervals); j++ {
				ret = append(ret, intervals[j])
			}
			return ret
		}
		// 此时 intervals[i][1]
		ret = append(ret, intervals[i])
	}
	return append(intervals, newInterval)
}
