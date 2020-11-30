package leetcode

type runecount struct {
	ch    rune
	count int
}

func reorganizeString(S string) string {
	table := make(map[rune]int)
	length := len(S)
	for _, ch := range S {
		table[ch]++
	}

	list := []runecount{}
	maxCount := 0
	for key, val := range table {
		list = append(list, runecount{key, val})
		maxCount = max(maxCount, val)
	}
	if maxCount > (length+1)/2 {
		return ""
	}
	ret := make([]rune, length)
	for i := 0; i < length; {
		ret[i] = list[0].ch
		list[0].count--
		if list[0].count == 0 {
			list = list[1:]
		}
		if (i == length-1 && length%2 == 1) || (i == length-2 && length%2 == 0) {
			i = 1
		} else {
			i += 2
		}
	}
	return string(ret)
}
