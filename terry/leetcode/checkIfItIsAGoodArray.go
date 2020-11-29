package leetcode

//斐蜀定理：
// 存在n个正整数 a1,a2,...,an，他们的最大公约数是d，那么肯定存在x1,x2,...,xn使得
// a1x1 + a2x2 + ... + anxn = d
// 所以我们只需要判断其最大公约数是否为1
func isGoodArray(nums []int) bool {
	ret := nums[0]
	for _, num := range nums {
		ret = gcd(ret, num)
	}
	return ret == 1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
