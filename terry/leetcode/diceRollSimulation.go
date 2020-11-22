package leetcode

const mod = 1e9 + 7

// 使用动态规划算法, dp[i][j]表示 第i(0 <= i < n)次投掷骰子是j(0 <= j < 6)的可能序列数
// 其中 dp[0][0], dp[0][1], ..., dp[0][5]都是1
//
func diceSimulator(n int, rollMax []int) int {

	dp := make([][6]int, n)
	for i := 0; i < 6; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < 6; j++ {
			ans := diceSum(dp[i-1])
			prej := i - rollMax[j] - 1
			if prej >= 0 { // i > rollMax[j]
				for k := 0; k < 6; k++ {
					if k != j {
						ans = (ans - dp[prej][k] + mod) % mod
					}
				}
			} else if prej == -1 { // i == rollMax[j]
				// 此时 i == rollMax[j]， 只需减去前i次（从0计数，0,1,...,i-1）都是j的情况
				ans--
			}
			dp[i][j] = ans % mod
		}
	}
	return diceSum(dp[n-1])
}

func diceSum(arr [6]int) (ret int) {
	for _, num := range arr {
		ret = (ret + num) % mod
	}
	return ret
}
