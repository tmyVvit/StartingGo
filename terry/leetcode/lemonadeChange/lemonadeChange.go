package main

import "fmt"

// 一杯柠檬汁 ¥5
// 每位顾客只会买一杯柠檬汁，会付我 $5, $10或者$20，我需要找零
// 最初我没有任何钱
func lemonadeChange(bills []int) bool {
	mymoney := make(map[int]int)
	for _, val := range bills {
		mymoney[val]++
		if val == 10 {
			if mymoney[5] == 0 {
				return false
			}
			mymoney[5]--
		} else if val == 20 {
			if mymoney[10] > 0 && mymoney[5] > 0 {
				mymoney[10]--
				mymoney[5]--
			} else if mymoney[5] >= 3 {
				mymoney[5] -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
}
