package leetcode

func fourSumCount(A []int, B []int, C []int, D []int) int {
	table := make(map[int]int)

	for _, a := range A {
		for _, b := range B {
			table[a+b]++
		}
	}

	ans := 0
	for _, c := range C {
		for _, d := range D {
			if val, ok := table[-c-d]; ok && val > 0 {
				ans += val
			}
		}
	}
	return ans
}
