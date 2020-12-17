package leetcode

// dp[i][j]
// i days , j 是否持有股
//  prices = [1, 3, 2, 8, 4, 9], fee = 2:
// dp[0][0] = 0  dp[0][1] = -price[0]
//    0                      -1
// dp[1][0] = max(dp[0][0], dp[0][1] + prices[1] - fee)  dp[1][1] = max(dp[0][0]-price[1], dp[0][1])
//    0           0          -1+3-2=0                      -1             0-3           -1
// dp[2][0] = max(dp[1][0], dp[1][1] + prices[2] - fee)  dp[2][1] = max(dp[1][0]-price[1], dp[1][1])
//     0             0           -1+2-2                     -1              0-2                -1
// dp[3][0] = max(dp[2][0], dp[2][1] + prices[3] - fee)  dp[3][1] = max(dp[2][0]-price[3], dp[2][1])
//     5             0          -1+8-2                     -1              0-8                -1
// dp[4][0] = max(dp[3][0], dp[3][1] + prices[4] - fee)  dp[4][1] = max(dp[3][0]-price[4], dp[3][1])
//     5             5          -1+4-2                    1             5-4                -1
// dp[5][0] = max(dp[4][0], dp[4][1] + prices[5] - fee)  dp[5][1] = max(dp[4][0]-price[5], dp[4][1])
//     8             5          1+9-2                     -1              5-9                1
func maxProfitWithTransactionFee(prices []int, fee int) int {
	if len(prices) == 0 || fee < 0 {
		return 0
	}
	days := len(prices)
	dp := make([][]int, days)
	dp[0] = []int{0, -prices[0]}
	for i := 1; i < days; i++ {
		tmp := []int{0, 0}
		tmp[0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		tmp[1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
		dp[i] = tmp
	}
	return max(dp[days-1][0], dp[days-1][1])
}

func maxProfitWithTransactionFeeI(prices []int, fee int) int {
	if len(prices) == 0 || fee < 0 {
		return 0
	}
	days := len(prices)
	pre := []int{0, -prices[0]}
	for i := 1; i < days; i++ {
		curr := []int{0, 0}
		curr[0] = max(pre[0], pre[1]+prices[i]-fee)
		curr[1] = max(pre[0]-prices[i], pre[1])
		pre = curr
	}
	return max(pre[0], pre[1])
}
