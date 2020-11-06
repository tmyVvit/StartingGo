package leetcode

import "sort"

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(a, b int) bool {
		tmpa := getNumberOf1Bits(arr[a])
		tmpb := getNumberOf1Bits(arr[b])
		if tmpa == tmpb {
			return arr[a] < arr[b]
		}
		return tmpa < tmpb

	})
	return arr
}

func getNumberOf1Bits(num int) (ret int) {
	for num != 0 {
		num = num & (num - 1)
		ret++
	}
	return ret
}
