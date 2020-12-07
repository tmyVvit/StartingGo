package leetcode

import "sort"

type numcount struct {
	num, count int
}

func minSetSize(arr []int) int {
	numCount := make(map[int]int)
	for _, num := range arr {
		numCount[num]++
	}
	numcounts := []numcount{}
	for key, val := range numCount {
		numcounts = append(numcounts, numcount{key, val})
	}
	sort.Slice(numcounts, func(a, b int) bool {
		return numcounts[a].count > numcounts[b].count
	})
	count := 0
	i := 0
	for ; count < len(arr)/2; i++ {
		count += numcounts[i].count
	}
	return i
}
