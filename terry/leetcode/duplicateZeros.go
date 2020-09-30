package leetcode

import "fmt"

// func main() {
// 	// duplicateZeros([]int{1, 0, 2, 3, 0, 4, 5, 0})
// 	duplicateZeros([]int{1, 5, 2, 0, 0, 0, 0, 0, 0})
// 	duplicateZeros([]int{1, 5, 2, 0, 6, 8, 0, 6, 0})
// }

func duplicateZeros(arr []int) {
	discard := 0
	length := len(arr)
	for i := 0; i < length-discard; i++ {
		if arr[i] == 0 {
			if i == length-discard-1 {
				arr[length-1] = 0
				length--
				break
			}
			discard++
		}
	}

	left := length - discard - 1
	for left >= 0 {
		if arr[left] == 0 {
			arr[left+discard] = 0
			discard--
			arr[left+discard] = 0
		} else {
			arr[left+discard] = arr[left]
		}
		left--
	}

	fmt.Println(arr)
}
