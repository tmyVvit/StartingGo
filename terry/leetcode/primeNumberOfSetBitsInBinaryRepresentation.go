package leetcode

// L, R in range [1, 10^6]
// -> L/R 小于20位

func countPrimeSetBits(L int, R int) (res int) {
	for i := L; i <= R; i++ {
		if isPrime(getBitCount(i)) {
			res++
		}
	}
	return res
}

func getBitCount(val int) (res int) {
	for val > 0 {
		if val&1 == 1 {
			res++
		}
		val >>= 1
	}
	return res
}

func isPrime(val int) bool {
	return val == 2 || val == 3 || val == 5 || val == 7 || val == 11 || val == 13 || val == 17 || val == 19
}
