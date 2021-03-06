package leetcode

import "strconv"

func compress(chars []byte) int {
	preCount, ind := 0, 0
	var pre byte
	for _, char := range chars {
		if preCount == 0 {
			chars[ind] = char
			ind++
			preCount, pre = 1, char
			continue
		}
		if char == pre {
			preCount++
		} else {
			if preCount > 1 {
				tmp := convertIntToByte(preCount)
				for i := len(tmp) - 1; i >= 0; i-- {
					chars[ind] = tmp[i]
					ind++
				}
			}
			chars[ind] = char
			ind++
			pre, preCount = char, 1
		}
	}
	if preCount > 1 {
		tmp := convertIntToByte(preCount)
		for i := len(tmp) - 1; i >= 0; i-- {
			chars[ind] = tmp[i]
			ind++
		}
	}
	return ind
}

func convertIntToByte(num int) (ret []byte) {
	for ; num > 0; num = num / 10 {
		ret = append(ret, byte(num%10+'0'))
	}
	return ret
}

func CompressI(chars []byte) int {
	pre, idx := 0, 0
	for i := 1; i <= len(chars); i++ {
		if i == len(chars) || chars[i-1] != chars[i] {
			chars[idx] = chars[pre]
			idx++
			count := i - pre
			if count > 1 {
				for _, char := range strconv.Itoa(count) {
					chars[idx] = byte(char)
					idx++
				}
			}
			pre = i
		}
	}
	return idx
}
