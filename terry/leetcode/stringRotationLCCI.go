package leetcode

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	length := len(s1)
	for i := 0; i < length; i++ {
		tmp := i
		for j := 0; j < length; j++ {
			if s1[j] == s2[tmp%length] {
				if j == length-1 {
					return true
				}
			}
			tmp++
		}
	}
	return false
}

// 如果s1 s2满足条件，那么(s1 + s1).contains(s2)
