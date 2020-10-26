package leetcode

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {
	copy := append([]int{}, nums...)
	sort.Ints(copy)
	mapp := make(map[int]int)
	sum, pre := 0, copy[0]
	for id, val := range copy {
		if pre == val {
			mapp[val] = sum
		} else {
			sum = id
			mapp[val] = sum
			pre = val
		}
	}
	res := []int{}
	for _, val := range nums {
		res = append(res, mapp[val])
	}
	return res
}

type pair struct {
	val, loc int
}

func smallerNumbersThanCurrentI(nums []int) []int {
	pairs := []pair{}
	for id, val := range nums {
		pairs = append(pairs, pair{val, id})
	}
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].val < pairs[b].val
	})
	pre, res := -1, make([]int, len(nums))
	for id, pair := range pairs {
		if pre == -1 || pair.val > pairs[id-1].val {
			pre = id
		}
		res[pair.loc] = pre
	}
	return res
}

func smallerNumbersThanCurrentII(nums []int) []int {
	// 考虑到 0 <= nums[i] <= 100
	// bigger[i]代表nums中小于i的数的个数
	bigger := make([]int, 101)
	for _, num := range nums {
		for j := num + 1; j < 101; j++ {
			// 所有大于num的均+1
			bigger[j]++
		}
	}
	res := make([]int, len(nums))
	for id, val := range nums {
		res[id] = bigger[val]
	}
	return res
}

func smallerNumbersThanCurrentIII(nums []int) []int {
	//使用dp
	//dp[0][i] 代表nums中i的个数
	//dp[1][i] 代表nums中小于i的数的个数
	// --> dp[1][i] = dp[1][i-1] + dp[0][i-1]
	dp := make([][101]int, 2)
	for _, val := range nums {
		dp[0][val]++
	}
	for i := 1; i < 101; i++ {
		dp[1][i] = dp[0][i-1] + dp[1][i-1]
	}
	ret := make([]int, len(nums))
	for id, val := range nums {
		ret[id] = dp[1][val]
	}
	return ret
}
