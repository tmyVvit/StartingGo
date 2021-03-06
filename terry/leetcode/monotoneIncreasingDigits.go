package leetcode

// MonotoneIncreasingDigitsI
func MonotoneIncreasingDitigsI(N int) int {
	if N < 10 {
		return N
	}
	// 123 -> []int{3,2,1}
	numberArr := int2arr(N)
	var i int
	// 找到不是递增的那一位
	for i = len(numberArr) - 1; i > 0; i-- {
		if numberArr[i-1] < numberArr[i] {
			break
		}
	}
	if i > 0 {
		for i < len(numberArr) && numberArr[i-1] < numberArr[i] {
			numberArr[i]--
			i++
		}
		// 因为上面循环最后有一个i++，所以找到减1那位数的后一位 需要 i-2
		i -= 2
		for i >= 0 {
			numberArr[i] = 9
			i--
		}
	}
	return arr2int(numberArr)
}

// MonotoneIncreasingDigits 返回小于等于给定数的最大整数，该整数满足各个位置上的数字单调递增
//回溯 + 递归
func MonotoneIncreasingDigits(N int) int {
	numberArr := int2arr(N)
	result := make([]int, len(numberArr))

	var findDigits func(nums []int, length int)
	findDigits = func(nums []int, length int) {
		if length < 1 {
			return
		}
		first := nums[length-1]
		// 判断当前位是否可以填对应的数字
		// 如果可以，则继续寻找下一位，否则当前位减1，并且后面的数全部为9
		if arrEachElementBiggerThan(nums, first, length) {
			result[length-1] = first
			findDigits(nums, length-1)
		} else {
			first--
			result[length-1] = first
			for i := 0; i < length-1; i++ {
				result[i] = 9
			}
		}
	}
	findDigits(numberArr, len(numberArr))
	return arr2int(result)
}

func arrEachElementBiggerThan(arr []int, first int, length int) bool {
	for i := length - 1; i >= 0; i-- {
		if arr[i] > first {
			return true
		} else if arr[i] < first {
			return false
		}
	}
	return true
}

func arr2int(arr []int) int {
	ret := 0
	for i := len(arr) - 1; i >= 0; i-- {
		ret = ret*10 + arr[i]
	}
	return ret
}

func int2arr(n int) []int {
	if n == 0 {
		return []int{0}
	}
	arr := []int{}
	for n > 0 {
		arr = append(arr, n%10)
		n /= 10
	}
	return arr
}
