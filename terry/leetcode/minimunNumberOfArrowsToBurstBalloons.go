package leetcode

import "sort"

// 首先按照区间左端点排序，然后计算交集
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(a, b int) bool {
		return points[a][0] < points[b][0]
	})
	rets := [][]int{[]int{points[0][0], points[0][1]}}
	for i := 1; i < len(points); i++ {
		var tmp []int
		for j := 0; j < len(rets); j++ {
			tmp = findSameSubs(points[i], rets[j])
			if tmp != nil {
				rets[j] = tmp
				break
			}
		}
		if tmp == nil {
			rets = append(rets, points[i])
		}
	}
	return len(rets)
}

func findSameSubs(point1 []int, point2 []int) []int {
	s1, e1, s2, e2 := point1[0], point1[1], point2[0], point2[1]
	if s1 > e2 || s2 > e1 {
		return nil
	}
	if s1 >= s2 {
		if e1 > e2 {
			return []int{s1, e2}
		}
		return []int{s1, e1}

	} else if s2 >= s1 {
		if e2 > e1 {
			return []int{s2, e1}
		}
		return []int{s2, e2}
	}
	return nil
}

// 将区间按照右端点排序
func findMinArrowShotsI(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(a, b int) bool {
		return points[a][1] < points[b][1]
	})
	ret := 1
	// 以最小的右端点为第一个射箭点
	right := points[0][1]
	for i := 1; i < len(points); i++ {
		// 只要当前区间的左端点在right左边，那么射在right点的箭肯定可以射穿这个区间
		if points[i][0] > right {
			ret++
			right = points[i][1]
		}
	}
	return ret
}
