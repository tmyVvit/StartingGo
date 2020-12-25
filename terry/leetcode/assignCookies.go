package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(a, b int) bool {
		return g[a] < g[b]
	})
	sort.Slice(s, func(a, b int) bool {
		return s[a] < s[b]
	})

	count, i, j := 0, 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			count++
			i++
			j++
		} else {
			j++
		}
	}
	return count
}
