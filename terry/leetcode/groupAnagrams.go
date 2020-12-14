package leetcode

import "sort"

func groupAnagrams(strs []string) (ret [][]string) {
	table := make(map[string][]string)
	for _, str := range strs {
		sorted := getSortAnagram(str)
		table[sorted] = append(table[sorted], str)
	}
	for _, v := range table {
		ret = append(ret, v)
	}
	return ret
}

func getSortAnagram(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(a, b int) bool {
		return runes[a] < runes[b]
	})
	return string(runes)
}
