package leetcode

func FirstUniqChar(s string) int {
	table := [26]int{}
	for idx, ch := range s {
		if table[ch-'a'] != 0 {
			table[ch-'a'] = -1
		} else {
			table[ch-'a'] = idx + 1
		}
	}
	first := -1
	for i := 0; i < 26; i++ {
		if table[i] > 0 {
			if first == -1 {
				first = table[i] - 1
			} else {
				first = min(first, table[i]-1)
			}
		}
	}
	return first
}
