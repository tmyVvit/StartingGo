package leetcode

func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}
	isPrimes := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrimes[i] = true
	}
	for i := 2; i*i < n; i++ {
		if !isPrimes[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			isPrimes[j] = false
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if isPrimes[i] {
			count++
		}
	}
	return count
}
