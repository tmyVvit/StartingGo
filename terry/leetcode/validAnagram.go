package leetcode

// 用两个数组分别记录每个字符串的字母出现个数，再比较
func isAnagram(s string, t string) bool {
	smap := make([]int, 26)
	tmap := make([]int, 26)
	for _, str := range s {
		smap[str-'a']++
	}
	for _, str := range t {
		tmap[str-'a']++
	}
	for i := 0; i < 26; i++ {
		if smap[i] != tmap[i] {
			return false
		}
	}
	return true
}

// 可以只用一个数组记录s的每个字符个数，再遍历t，没出现一个便减去1，如果smap[i] < 0那么就不是
func isAnagramI(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	smap := make([]int, 26)
	for _, str := range s {
		smap[str-'a']++
	}
	for _, str := range t {
		smap[str-'a']--
		if smap[str-'a'] < 0 {
			return false
		}
	}
	return true
}
