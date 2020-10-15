package leetcode

import "math"

func maxProfitII(prices []int) int {
	sum, max, prev := 0, 0, 0
	minPrice := math.MaxInt32
	for _, price := range prices {
		// 只要股票价格下降，那么就要在下降之前将股票卖出
		if price < prev {
			sum += max
			max = 0
			minPrice = price
		} else if price < minPrice {
			minPrice = price
		} else if max < price-minPrice {
			max = price - minPrice
		}
		prev = price
	}
	sum += max
	return sum
}
