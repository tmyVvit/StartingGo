package leetcode

//老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
// 你需要按照以下要求，帮助老师给这些孩子分发糖果：
// 	 1. 每个孩子至少分配到 1 个糖果。
// 	 2. 相邻的孩子中，评分高的孩子必须获得更多的糖果。
// 那么这样下来，老师至少需要准备多少颗糖果呢？

// 暴力求解，每次循环检查是否符合要求，不符合就增加糖果
// 只要有增加糖果，就需要进行下一次循环检查
func Candy(ratings []int) int {
	needRecalculate := true
	count := len(ratings)
	candies := make([]int, count)
	for needRecalculate {
		needRecalculate = false
		for i := 1; i < count; i++ {
			if ratings[i] == ratings[i-1] {
				continue
			} else if ratings[i] > ratings[i-1] {
				if candies[i] <= candies[i-1] {
					candies[i] = candies[i-1] + 1
					needRecalculate = true
				}
			} else {
				if candies[i] >= candies[i-1] {
					candies[i-1] = candies[i] + 1
					needRecalculate = true
				}
			}
		}
	}
	sum := 0
	for _, num := range candies {
		sum += num
	}
	return sum + count
}

// 分别从左到右，从右到左两次循环遍历
// 如果当前数比之前的要大，则加1，否则置为1
// 最终结果取两次遍历的最大值
func candyI(ratings []int) int {
	left := make([]int, len(ratings))
	for i := 0; i < len(left); i++ {
		if i == 0 || ratings[i] <= ratings[i-1] {
			left[i] = 1
		} else {
			left[i] = left[i-1] + 1
		}
	}
	right := 1
	ret := 0
	for i := len(left) - 1; i >= 0; i-- {
		if i == len(left)-1 || ratings[i] <= ratings[i+1] {
			right = 1
		} else {
			right++
		}
		ret += max(left[i], right)
	}
	return ret
}
