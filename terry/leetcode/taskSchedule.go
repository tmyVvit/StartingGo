package leetcode

import "sort"

// https://leetcode-cn.com/problems/task-scheduler/solution/tong-zi-by-popopop/
func leastInterval(tasks []byte, n int) int {
	table := make([]int, 26)
	for _, task := range tasks {
		table[task-'A']++
	}
	sort.Slice(table, func(a, b int) bool {
		return table[a] < table[b]
	})
	maxVal, idles := table[25]-1, (table[25]-1)*n
	for i := 24; i >= 0; i-- {
		idles -= min(maxVal, table[i])
	}
	if idles > 0 {
		return idles + len(tasks)
	}
	return len(tasks)
}

// 设 n+1个任务为一轮，每次取任务数最多的 n+1 个任务
// 每次取完任务后 重新排序
func leastIntervalI(tasks []byte, n int) int {
	table := make([]int, 26)
	for _, task := range tasks {
		table[task-'A']++
	}
	sort.Slice(table, func(a, b int) bool {
		return table[a] < table[b]
	})
	ret := 0
	for table[25] > 0 {
		for i := 0; i < n+1; i++ {
			if table[25] == 0 {
				break
			}
			if i < 26 && table[25-i] > 0 {
				table[25-i]--
			}
			ret++
		}
		sort.Slice(table, func(a, b int) bool {
			return table[a] < table[b]
		})
	}
	return ret
}
