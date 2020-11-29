package leetcode

// 1. 暴力求解是否回文
func shortestPalindrome(s string) string {
	ch := []rune(s)
	i := 0
	for ; i < len(ch); i++ {
		tmp := ch[:len(ch)-i]
		if isPalidrome(tmp) {
			break
		}
	}
	ret := []rune{}
	for j := 0; j < i; j++ {
		ret = append(ret, ch[len(ch)-j])
	}
	ret = append(ret, ch...)
	return string(ret)
}

func isPalidrome(ch []rune) bool {
	length := len(ch)
	for i := (length - 1) / 2; i >= 0; i-- {
		if ch[i] != ch[length-1-i] {
			return false
		}
	}
	return true
}

// 考虑KMP算法的思路
// 我们需要找到字符串最长的一个回文字符子串，那么我们可以将字符串和它的翻转字符串拼接起来
// 例如："abcda" "adcba" 就拼接成了 "abcda#adcba", 那么就可以看成是找到这个拼接字符串的最长公共前后缀
// 也就是要求KMP算法中的next数组的next[n]
func shortestPalindromeI(s string) string {
	ch := []rune(s)
	length := len(ch)
	reverse := make([]rune, length)
	for id, c := range ch {
		reverse[length-1-id] = c
	}
	maxPalindromeLength := getLatestNext(append(append(ch, '#'), reverse...))
	return string(append(reverse[:length-maxPalindromeLength], ch...))
}

func getLatestNext(ch []rune) int {
	length := len(ch)
	// next[i] 代表 ch[:i-1]的最长公共前后缀的长度
	next := make([]int, length+1)
	next[0] = -1
	next[1] = 0
	i, k := 2, 0
	for i <= length {
		if k == -1 || ch[i-1] == ch[k] {
			next[i] = k + 1
			k++
			i++
		} else {
			k = next[k]
		}
	}
	return next[length]
}
