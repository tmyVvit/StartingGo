package leetcode

// 使用二分法
func firstBadVersion(n int) int {
	start, end := 1, n
	for start < end {
		mid := (end + start) / 2
		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

// this should be given by leetcode
func isBadVersion(version int) bool {
	if version >= 3 {
		return true
	}
	return false
}
