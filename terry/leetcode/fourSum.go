package leetcode

import (
	"fmt"
	"sort"
)

type twoSum struct {
	sum  int
	a, b int
}

// 循环计算两个数的和 a+b， 将 a+b放入map中，并检查 target - a - b是否在map中
// 如果在，则加入答案中（注意重复）
func fourSum(nums []int, target int) (ret [][]int) {
	if len(nums) == 0 {
		return [][]int{}
	}
	twosums := make(map[int][]twoSum)
	// 是否已经加入重复的值
	used := make(map[string]bool)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			twosums[sum] = append(twosums[sum], twoSum{sum: sum, a: i, b: j})
			negSum := target - nums[i] - nums[j]
			for _, sums := range twosums[negSum] {
				if i != sums.a && i != sums.b && j != sums.a && j != sums.b && used[buildSumKey(i, j, sums.a, sums.b)] {
					ret = append(ret, []int{nums[i], nums[j], nums[sums.a], nums[sums.b]})
				}
			}
		}
	}
	return ret
}

func buildSumKey(a, b, c, d int) string {
	nums := []int{a, b, c, d}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return fmt.Sprintf("%03d%03d%03d%03d", nums[0], nums[1], nums[2], nums[3])
}
