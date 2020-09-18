package permutations2

import "sort"

func permuteUnique(nums []int) (res [][]int) {
	sort.Ints(nums)
	used := make([]bool, len(nums))
	curr := make([]int, len(nums))

	var traceback func(idx int, curr []int)
	traceback = func(idx int, curr []int) {
		if idx == len(nums) {
			res = append(res, append([]int{}, curr...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}
			used[i] = true
			curr[idx] = nums[i]
			traceback(idx+1, curr)
			curr[idx] = 0
			used[i] = false
		}
	}
	traceback(0, curr)
	return res
}

func traceback(nums []int, idx int, curr []int, used []bool, res [][]int) {
	if idx == len(nums) {
		res = append(res, append([]int{}, curr...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
			continue
		}
		used[i] = true
		curr[idx] = nums[i]
		traceback(nums, idx+1, curr, used, res)
		curr[idx] = 0
		used[i] = false
	}
}
