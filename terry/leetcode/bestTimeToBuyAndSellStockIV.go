package leetcode

// dp[i][j][k]
//   j = 0  没有股票   1  有股票
//   k = 已经卖出的次数
// j = 0 时：
//    k = 0: dp[i][0][0] = dp[i-1][0][0]
//    k = 1: dp[i][0][1] = max(dp[i-1][0][1], dp[i-1][1][0] + price[i])
//    k = 2: dp[i][0][2] = max(dp[i-1][0][2], dp[i-1][1][1] + price[i])
// j = 1:
//    k = 0: dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][0][0] - price[i])    /* 前一天已经持有，或者前一天未持有 当天刚买入 */
//    k = 1: dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][1] - price[i])
//    k = 2: dp[i][1][2]   因为只能交易两次，所以不可能出现已经卖出两次并且手中还持有股票的情况

// 上面是 best time to buy and sell stock III 最多只能2次交易的动态方程
// 我们可以参考上面的解法：
// j = 0 时：
//    k = 0: dp[i][0][0] = dp[i-1][0][0]
//    k > 0: dp[i][0][k] = max(dp[i-1][0][k], dp[i-1][1][k-1] + price[i])
// j = 1:
//    k = 0: dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][0][0] - price[i])    /* 前一天已经持有，或者前一天未持有 当天刚买入 */
//    k > 0: dp[i][1][k] = max(dp[i-1][1][k], dp[i-1][0][k] - price[i])

// MaxProfitIV 返回最多交易k次的最大利润
func MaxProfitIV(k int, prices []int) int {
	days := len(prices)
	if days < 2 {
		return 0
	}
	// 因为转移方程只需要用到前一天的数值，所以我们只需要保存前一天的数据
	predp := [2][]int{}
	curdp := [2][]int{}
	day0 := make([]int, k+1)
	predp[0] = day0
	dayk := make([]int, k+1)
	for j := 0; j <= k; j++ {
		dayk[j] = -prices[0]
	}
	predp[1] = dayk
	for i := 1; i < days; i++ {
		dayi := make([]int, k+1)
		for j := 1; j <= k; j++ {
			dayi[j] = max(predp[0][j], predp[1][j-1]+prices[i])
		}
		curdp[0] = dayi
		dayi = make([]int, k+1)
		for j := 0; j < k; j++ {
			dayi[j] = max(predp[1][j], predp[0][j]-prices[i])
		}
		curdp[1] = dayi
		predp = curdp
		curdp = [2][]int{}
	}
	maxp := 0
	for i := 0; i <= k; i++ {
		maxp = max(maxp, predp[0][i])
	}
	return maxp
}
