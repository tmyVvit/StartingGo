package leetcode

import (
	"math"
)

func commonChars(A []string) (res []string) {
	minFreq := [26]int{}
	for i := 0; i < 26; i++ {
		minFreq[i] = math.MaxInt16
	}
	for _, str := range A {
		freq := [26]int{}
		for _, char := range str {
			freq[char-'a']++
		}
		for i := 0; i < 26; i++ {
			if freq[i] < minFreq[i] {
				minFreq[i] = freq[i]
			}
		}
	}
	for i := 0; i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			res = append(res, string(i+'a'))
		}
	}
	return res
}
