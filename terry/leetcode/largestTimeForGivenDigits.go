package leetcode

import "fmt"

// 贪心 + 回溯
func largestTimeFromDigits(arr []int) string {
	limits := []int{2, 9, 5, 9}
	table := make([]int, 10)
	for _, num := range arr {
		table[num]++
	}
	ret := []int{}
	var dfs func(nums []int) bool
	dfs = func(nums []int) bool {
		if len(nums) == 4 {
			return true
		}
		limit := limits[len(nums)]
		// 当第一位是2时，最大是3 （小时小于24）
		if limit == 9 && len(nums) == 1 && nums[0] == 2 {
			limit = 3
		}
		for i := limit; i >= 0; i-- {
			if table[i] == 0 {
				continue
			}
			table[i]--
			ret = append(ret, i)
			if dfs(ret) {
				return true
			}
			table[i]++
			ret = ret[:len(ret)-1]
		}
		return false
	}
	if dfs(ret) {
		return fmt.Sprintf("%d%d:%d%d", ret[0], ret[1], ret[2], ret[3])
	}
	return ""
}

// 由于数组只有4位，所以可以直接暴力求解
func largestTimeFromDigitsI(arr []int) string {
	ret := -1
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j {
				continue
			}
			for k := 0; k < 4; k++ {
				if i == k || j == k {
					continue
				}
				l := 6 - i - j - k
				hour := 10*arr[i] + arr[j]
				minute := 10*arr[k] + arr[l]
				if hour < 24 && minute < 59 {
					ret = max(ret, 60*hour+minute)
				}
			}
		}
	}
	if ret >= 0 {
		return fmt.Sprintf("%02d:%02d", ret/3600, ret/60)
	}
	return ""
}
