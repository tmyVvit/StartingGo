package leetcode

import (
	"container/heap"
	"sort"
)

func kClosest(points [][]int, K int) [][]int {
	if K >= len(points) {
		return points
	}
	sort.Slice(points, func(a, b int) bool {
		return (square(points[a][0]) + square(points[a][1])) < (square(points[b][0]) + square(points[b][1]))
	})
	heap.Init()
	return points[:K]
}

func square(a int) int {
	return a * a
}
