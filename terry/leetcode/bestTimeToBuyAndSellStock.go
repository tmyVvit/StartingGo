package leetcode

import "math"

//  profit[i]  第i天最好收益
//  min[i] 前i天最低价格
//  min[i+1] = Min(min[i], prices[i+1])
//  profit[i+1] = 1. prices[i+1] > min[i] : Max(profit[i], prices[i+1] - min[i])
//				  2. prices[i+1] <= min[i]: profit[i]
func maxProfit(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	profit := make([]int, length)
	min := make([]int, length)
	min[0] = prices[0]
	for i := 1; i < length; i++ {
		if prices[i] > min[i-1] {
			profit[i] = Max(profit[i-1], prices[i]-min[i-1])
		} else {
			profit[i] = profit[i-1]
		}
		min[i] = Min(min[i-1], prices[i])
	}
	return profit[length-1]
}

func maxProfit2(prices []int) int {
	maxProfit := 0
	minPrices := math.MaxInt32
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrices {
			minPrices = prices[i]
		} else if maxProfit < prices[i]-minPrices {
			maxProfit = prices[i] - minPrices
		}
	}
	return maxProfit
}
