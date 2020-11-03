package leetcode

import "math"

// WordBreakII 理论上是可行的
// 但是在遇到wordDict全由a组成的字符串["a", "aa", "aaa", ...]，并且string是"aaaaaaaaaaaaa..."时会超时
func WordBreakII(s string, wordDict []string) (ret []string) {
	dict := make(map[string]bool)
	minlen, maxlen := math.MaxInt32, 0
	for _, word := range wordDict {
		dict[word] = true
		minlen = min(len(word), minlen)
		maxlen = max(len(word), maxlen)
	}
	var breakword func(str string, words []string)
	breakword = func(str string, words []string) {
		if len(str) < minlen {
			return
		}
		for i := minlen; i <= maxlen && i <= len(str); i++ {
			if dict[str[:i]] {
				if i == len(str) {
					// 此时说明已经将string找完了，
					ret = append(ret, buildSentence(append(words, str[:i])))
					return
				}
				breakword(str[i:], append(words, str[:i]))
			}
		}
	}
	breakword(s, nil)
	return ret
}

func buildSentence(words []string) (ret string) {
	for _, word := range words {
		ret += (word + " ")
	}
	return ret[:len(ret)-1]
}

func buildSentences(wordList [][]string) (ret []string) {
	for _, words := range wordList {
		ret = append(ret, buildSentence(words))
	}
	return ret
}

// WordBreakII2 和 WordBreakII思想相同，但是增加了一个map来记录中间数据，避免重复计算
func WordBreakII2(s string, wordDict []string) []string {
	dict := make(map[string]bool)
	minlen, maxlen := math.MaxInt32, 0
	for _, word := range wordDict {
		dict[word] = true
		minlen = min(len(word), minlen)
		maxlen = max(len(word), maxlen)
	}
	// table[i]记录 s[i:] 可以组成句子的单词集合
	table := make(map[int][][]string)
	var breakword func(s string, index int, table map[int][][]string) [][]string
	breakword = func(s string, index int, table map[int][][]string) [][]string {
		if table[index] == nil {
			wordset := [][]string{}
			if index == len(s) {
				wordset = append(wordset, []string{})
			}
			for i := index + minlen; i <= len(s) && i <= index+maxlen; i++ {
				tmp := s[index:i]
				if dict[tmp] {
					subbreaks := breakword(s, i, table)
					for _, words := range subbreaks {
						addme := []string{}
						addme = append(addme, tmp)
						addme = append(addme, words...)
						wordset = append(wordset, addme)
					}
				}
			}
			table[index] = wordset
		}
		return table[index]
	}
	wordsList := breakword(s, 0, table)
	return buildSentences(wordsList)
}
