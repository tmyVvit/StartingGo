package leetcode

import "sort"

// 当 A + B > C 时可以组成一个三角形
func largestPerimeter(A []int) int {
	sort.Slice(A, func(a, b int) bool {
		return A[a] < A[b]
	})
	for i := len(A) - 1; i >= 2; i-- {
		if A[i] < A[i-1]+A[i-2] {
			return A[i] + A[i-1] + A[i-2]
		}
	}
	return 0
}
