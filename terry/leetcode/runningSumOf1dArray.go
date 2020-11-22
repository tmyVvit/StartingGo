package leetcode

func runningSum(nums []int) []int {
	sums := make([]int, len(nums))
	pre := 0
	for i := 0; i < len(nums); i++ {
		sums[i] = pre + nums[i]
		pre = sums[i]
	}
	return sums
}

//或者可以直接再原本的nums数组上更改
func runningSumI(nums []int) []int {
	if len(nums) < 1 {
		return nums
	}
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	return nums
}
