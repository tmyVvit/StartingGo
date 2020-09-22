package singleNumberIII

// Given an integer array nums, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once. You can return the answer in any order.
// Follow up: Your algorithm should run in linear runtime complexity. Could you implement it using only constant space complexity?

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/single-number-iii

/*
1. 可以直接遍历两边将出现一次的数字找出，但是这样时间复杂度和空间复杂度都是 O(N)
2. 使用异或操作，相同的数字的异或为0，所以数组中全部数字的异或结果就是只出现一次的两个数字的异或结果，
   然后找到一个不一样的位，继续和数组所有的数字做异或操作，就可以找到一个数字

*/

func singleNumber(nums []int) []int {
	var mask int
	for _, val := range nums {
		mask ^= val
	}
	diff := mask & (-mask) // a & (-a) 可以得到a的最低非零位
	var x int

	for _, val := range nums {
		if diff&val > 0 {
			x ^= val
		}
	}
	return []int{x, mask ^ x}
}
