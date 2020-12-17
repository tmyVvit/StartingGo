package leetcode

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	word2byte := make(map[string]byte)
	byte2word := make(map[byte]string)
	for i, word := range words {
		ch := pattern[i]
		cword := byte2word[ch]
		cbyte := word2byte[word]
		if cword != "" && cword != word || cbyte != 0 && cbyte != ch {
			return false
		}
		byte2word[ch] = word
		word2byte[word] = ch
	}
	return true
}
