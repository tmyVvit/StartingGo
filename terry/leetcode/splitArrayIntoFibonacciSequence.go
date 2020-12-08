package leetcode

import (
	"errors"
	"math"
	"strconv"
)

// 固定前两个数，暴力搜索
func SplitIntoFibonacci(S string) []int {
	if len(S) <= 2 {
		return []int{}
	}
	// 注意拆解出来的每个数都不大于 2^31 - 1
	const maxnum = math.MaxInt32
	var ret []int
	for first := 1; first <= len(S)/2; first++ {
		firstNum, err := getNumFromString(S, 0, first)
		if err != nil {
			continue
		}
		// 如果第一个值大于最大值，则不需要再找了
		if firstNum > maxnum {
			break
		}
		for second := 1; first+second < len(S); second++ {
			pre, err := getNumFromString(S, first, second)
			if err != nil {
				continue
			}
			// 如果第二个值大于最大值，则退出当前循环，重新寻找第一个值
			if pre > maxnum {
				break
			}
			prepre := firstNum
			ret = []int{prepre, pre}
			offset := first + second
			for {
				curr := prepre + pre
				if curr > maxnum {
					break
				}
				currLengh := getNumLength(curr)
				numFromStr, err := getNumFromString(S, offset, currLengh)
				if err != nil || curr != numFromStr {
					break
				}
				ret = append(ret, curr)
				offset += currLengh
				if offset == len(S) {
					return ret
				}
				prepre = pre
				pre = curr
			}
		}
	}
	return []int{}
}

// 从字符串中拿到数字
func getNumFromString(S string, offset int, length int) (int, error) {
	if offset < 0 || length <= 0 || offset+length > len(S) {
		return 0, errors.New("invalid input")
	}
	numarr := ([]rune)(S)
	// 数字开头不能有多余的0
	if numarr[offset] == '0' && length > 1 {
		return 0, errors.New("cannot have extra leading zeros")
	}
	number := string(numarr[offset : length+offset])
	return strconv.Atoi(number)
}

func getNumLength(num int) int {
	// 注意当num==0时，是占一位的
	if num == 0 {
		return 1
	}
	count := 0
	for num > 0 {
		num /= 10
		count++
	}
	return count
}
