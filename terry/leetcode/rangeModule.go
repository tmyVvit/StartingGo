package leetcode

// RangeModule is a struct that tracks range of numbers
type RangeModule struct {
	arr [][]int
}

// NewRangeModule init a new RangeModule
func NewRangeModule() RangeModule {
	return RangeModule{arr: [][]int{}}
}

// AddRange add [left, right) to range
func (rm *RangeModule) AddRange(left int, right int) {
	rm.arr = insertRange(rm.arr, []int{left, right})
}

// QueryRange Returns true if and only if every real number in the interval [left, right) is currently being tracked
func (rm *RangeModule) QueryRange(left int, right int) bool {
	for _, r := range rm.arr {
		if r[0] <= left && r[1] >= right {
			return true
		}
	}
	return false
}

// RemoveRange Stops tracking every real number currently being tracked in the interval [left, right).
func (rm *RangeModule) RemoveRange(left int, right int) {
	newArr := [][]int{}
	for _, r := range rm.arr {
		start, end := r[0], r[1]
		if start < right && end > left {
			splits := minusRange([]int{start, end}, []int{left, right})
			if len(splits) > 0 {
				newArr = append(newArr, splits...)
			}
		} else {
			newArr = append(newArr, r)
		}
	}
	rm.arr = newArr
}

func minusRange(array []int, del []int) (ret [][]int) {
	if del[0] <= array[0] && del[1] >= array[1] {
		return ret
	} else if del[0] <= array[0] && del[1] < array[1] && del[1] > array[0] {
		return append(ret, []int{del[1], array[1]})
	} else if del[0] > array[0] && del[1] > array[1] {
		return append(ret, []int{array[0], del[0]})
	} else if del[0] > array[0] && del[1] < array[1] {
		return append(ret, []int{array[0], del[0]}, []int{del[1], array[1]})
	}
	return ret
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := NewRangeModule();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */

func insertRange(intervals [][]int, newInterval []int) (ret [][]int) {
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
