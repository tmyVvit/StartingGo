package leetcode

import "math"

func maxNumberOfBalloons(text string) int {
	table := make(map[rune]int)
	for _, ch := range text {
		table[rune(ch)] = table[rune(ch)] + 1
	}

	return minList(table['b'], table['a'], table['l']/2, table['o']/2, table['n'])

}

func minList(a ...int) int {
	minNum := math.MaxInt32
	for _, val := range a {
		if minNum > val {
			minNum = val
		}
	}
	return minNum
}
