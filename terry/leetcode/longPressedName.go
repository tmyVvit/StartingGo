package leetcode

func isLongPressedName(name string, typed string) bool {
	if len(name) > len(typed) {
		return false
	}

	i, j := 0, 0

	for i < len(name) && j < len(typed) {
		if name[i] != typed[j] {
			return false
		}
		tmpCh := name[i]
		numName, numTyped := 0, 0
		for ; i < len(name) && name[i] == tmpCh; i++ {
			numName++
		}
		for ; j < len(typed) && typed[j] == tmpCh; j++ {
			numTyped++
		}
		if numTyped < numName {
			return false
		}
	}
	return i == len(name) && j == len(typed)
}
