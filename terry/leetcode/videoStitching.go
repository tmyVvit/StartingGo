package leetcode

// VideoStitching 贪心算法
func VideoStitching(clips [][]int, T int) int {
	maxn := make([]int, T)
	for _, clip := range clips {
		// 子区间又端点的值可能大于T
		if clip[0] < T {
			// maxn[i] 代表包含i的子区间的最大可达右端点
			maxn[clip[0]] = max(maxn[clip[0]], clip[1])
		}
	}
	pre, last, ret := 0, 0, 0
	for i := 0; i < T; i++ {
		if last < maxn[i] {
			last = maxn[i]
		}
		if i == last {
			// 说明此时中间已经断档了， last 到 last + 1无法被覆盖
			return -1
		}
		if i == pre {
			// 上一个子区间已经走完
			// 继续下一个子区间，取右端点最大的子区间（贪心）
			pre = last
			ret++
		}
	}
	return ret
}
