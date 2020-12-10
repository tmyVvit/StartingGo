package leetcode

// 使用动态规划 dp[i][j] 表示从起点到(i,j)处的路径数
// dp[0][0] = 1 , dp[i][j] = dp[i-1][j] + dp[i][j-1]
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1]
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
		}
	}
	return dp[m-1][n-1]
}

// 可以只使用一个一维数组实现
func uniquePathsI(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}
