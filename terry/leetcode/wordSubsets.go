package leetcode

// 首先将B中所有string合并为一个
// 然后遍历A，判断是否是A的subset
func wordSubsets(A []string, B []string) (ret []string) {
	bTable := make(map[rune]int)
	for _, b := range B {
		tmp := buildCharTable(b)
		for key, val := range tmp {
			bTable[key] = max(bTable[key], val)
		}
	}

	for _, a := range A {
		aTable := buildCharTable(a)
		isSubset := true
		for key, val := range bTable {
			if val > aTable[key] {
				isSubset = false
				break
			}
		}
		if isSubset {
			ret = append(ret, a)
		}
	}
	return ret
}

func buildCharTable(a string) map[rune]int {
	tmp := make(map[rune]int)
	for _, ch := range a {
		tmp[ch] = tmp[ch] + 1
	}
	return tmp
}

// 和上面思想相同，但是不使用map实现
// 因为小写字母只有26位，所以可以使用数组实现
func wordSubsetsI(A []string, B []string) (ret []string) {
	bmax := make([]int, 26)
	for _, b := range B {
		tmp := countLetters(b)
		for i := 0; i < 26; i++ {
			bmax[i] = max(bmax[i], tmp[i])
		}
	}

	for _, a := range A {
		amax := countLetters(a)
		isSub := true
		for i := 0; i < 26; i++ {
			if bmax[i] > amax[i] {
				isSub = false
				break
			}
		}
		if isSub {
			ret = append(ret, a)
		}
	}
	return ret
}

func countLetters(a string) []int {
	count := make([]int, 26)
	for _, ch := range a {
		count[ch-'a']++
	}
	return count
}
