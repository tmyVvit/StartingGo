package leetcode

// 暴力遍历
func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	time := 0
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i] >= timeSeries[i-1]+duration {
			time += duration
		} else {
			time += (timeSeries[i] - timeSeries[i-1])
		}
	}
	return time + duration
}

// pre + duration
