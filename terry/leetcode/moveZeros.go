package leetcode

// MoveZeroes 在原数组的基础上，将所有的0移至数组末尾，并且非零数的顺序不变
// 使用一个数组记录当前位置之前的0的个数
func MoveZeroes(nums []int) {
	preZeros := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		preZeros[i] = preZeros[i-1]
		if nums[i-1] == 0 {
			preZeros[i] = preZeros[i] + 1
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && preZeros[i] > 0 {
			nums[i-preZeros[i]] = nums[i]
			nums[i] = 0
		}
	}
}

// MoveZeroesI 使用变量newIndex记录移动后的非零数的下标
func MoveZeroesI(nums []int) {
	newIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[newIndex] = nums[i]
			if newIndex != i {
				nums[i] = 0
			}
			newIndex++
		}
	}
}

// MoveZeroesII 与第一个解法思想相同，只用一个变量存储当前位置之前的零的个数
func MoveZeroesII(nums []int) {
	snowballSize := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			snowballSize++
		} else {
			nums[i], nums[i-snowballSize] = nums[i-snowballSize], nums[i]
		}
	}
}
