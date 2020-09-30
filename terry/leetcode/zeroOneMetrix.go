package leetcode

import "math"

// 01 matrix

func updateMatrix(matrix [][]int) [][]int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	x, y := len(matrix), len(matrix[0])
	dist := make([][]int, x)
	for i := 0; i < x; i++ {
		dist[i] = make([]int, y)
	}
	queue := [][]int{}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	// BFS 广度优先搜索
	for len(queue) > 0 {
		tmp := queue
		queue = nil
		for _, val := range tmp {
			m, n := val[0], val[1]
			for _, dir := range dirs {
				mi, ni := m+dir[0], n+dir[1]
				if mi < x && mi >= 0 && ni < y && ni >= 0 && (matrix[mi][ni] > 0 && dist[mi][ni] == 0) {
					dist[mi][ni] = dist[m][n] + 1
					queue = append(queue, []int{mi, ni})
				}
			}
		}
	}
	return dist
}

// 也可以考虑使用 动态规划
func updateMatrix2(matrix [][]int) [][]int {
	x, y := len(matrix), len(matrix[0])
	dp := make([][]int, x)
	for i := 0; i < x; i++ {
		dp[i] = make([]int, y)
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if matrix[i][j] == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = math.MaxInt32
			}
		}
	}

	// 从 左上 到 右下
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if i-1 >= 0 {
				dp[i][j] = Min(dp[i][j], dp[i-1][j]+1)
			}
			if j-1 >= 0 {
				dp[i][j] = Min(dp[i][j], dp[i][j-1]+1)
			}
		}
	}

	// 从 右下 到 左上
	for i := x - 1; i >= 0; i-- {
		for j := y - 1; j >= 0; j-- {
			if i+1 < x {
				dp[i][j] = Min(dp[i][j], dp[i+1][j]+1)
			}
			if j+1 < y {
				dp[i][j] = Min(dp[i][j], dp[i][j+1]+1)
			}
		}
	}
	return dp
}
