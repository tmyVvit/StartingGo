package leetcode

func uniqueOccurrences(arr []int) bool {
	table := make(map[int]int)
	for _, val := range arr {
		table[val] = table[val] + 1
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] != arr[j] && table[arr[i]] == table[arr[j]] {
				return false
			}
		}
	}
	return true
}
