package leetcode

func trailingZeroes(n int) int {
	five, two := 0, 0
	for i := 2; i <= n; i++ {
		for tmp := i; tmp%2 == 0; tmp /= 5 {
			two++
		}
		for tmp := i; tmp%5 == 0; tmp /= 5 {
			five++
		}
	}
	return min(two, five)
}

func trailingZeroesI(n int) int {
	five := 0
	for i := 5; i <= n; i += 5 {
		for tmp := i; tmp%5 == 0; tmp = tmp / 5 {
			five++
		}

	}
	return five
}
