package leetcode

import "math"

// MaxProfitIII 返回最多两次交易的最大利润
// 从左到右，从右到左分别计算一次最大利润
// left[i] 代表 0 -> i 之间的最大利润
// right[i] 代表 i -> n 之间的最大利润
func MaxProfitIII(prices []int) int {
	days := len(prices)
	if days <= 1 {
		return 0
	}
	left, right := make([]int, days), make([]int, days)
	left[0] = 0
	right[days-1] = 0
	minPrice := prices[0]
	for i := 1; i < days; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		left[i] = max(left[i-1], prices[i]-minPrice)
	}

	maxPrice := prices[days-1]
	maxProfit := 0
	for i := days - 2; i >= 0; i-- {
		if maxPrice < prices[i] {
			maxPrice = prices[i]
		}
		right[i] = max(right[i+1], maxPrice-prices[i])
		maxProfit = max(maxProfit, left[i]+right[i])
	}
	return maxProfit
}

//  ？？
func maxProfitIII(prices []int) int {
	cost1, cost2 := math.MaxInt32, math.MaxInt32
	profit1, profit2 := 0, 0
	for _, price := range prices {
		cost1 = min(cost1, price)
		profit1 = max(profit1, price-cost1)

		cost2 = min(cost2, price-profit1)
		profit2 = max(profit2, price-cost2)
	}
	return profit2
}

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
func maxProfitIII2(prices []int) int {
	days := len(prices)
	if days < 2 {
		return 0
	}
	dp := make([][2][3]int, days)
	// 第一天不持有，不管卖出多少次，都是0
	// 第一天持有，不管卖出多少次，都是 -prices[0], 因为不可能卖出，如果持有则说明有买入
	dp0 := [2][3]int{{0, 0, 0}, {-prices[0], -prices[0], -prices[0]}}
	dp = append(dp, dp0)

	for i := 0; i < days; i++ {
		dpi := [2][3]int{}
		dpi[0][1] = max(dp[i-1][0][1], dp[i-1][1][0]+prices[i])
		dpi[0][2] = max(dp[i-1][0][2], dp[i-1][1][1]+prices[i])
		dpi[1][0] = max(dp[i-1][1][0], dp[i-1][0][0]-prices[i])
		dpi[1][1] = max(dp[i-1][1][1], dp[i-1][0][1]-prices[i])
		dp = append(dp, dpi)
	}
	return dp[days-1][0][2]

}
