package leetcode

import "math"

// LeastOpsExpressTarget
// 我们将x^e看作一个表达式，那么最后的总的表达式肯定是 x^e1 (+/-) x^e2 (+/-) .... x^en
// 我们在最前面加一个加号： (+) x^e1 (+/-) x^e2 (+/-) .... x^en
// 表达式中的每一块都是 x的幂
// 我们只需计算出所有块所需运算符的合最后再减去1 即可
func LeastOpsExpressTarget(x int, target int) int {
	table := make(map[int]int)
	var find func(curr int) int
	find = func(curr int) int {
		if val, ok := table[curr]; ok {
			return val
		}
		ans := 0
		if curr == 0 {
			ans = 0
		} else if curr < x {
			// 当curr < x时，有两种情况，要么是全部1相加，或者是x减去x-curr个1
			// 比如 curr = 2, x = 3:
			// 2*curr       : + 3/3 + 3/3
			// 2*(x-curr)+1 : + 3 - 3/3
			ans = min(2*curr, 2*(x-curr)+1)
		} else {
			// 计算p，使的x^p是小于curr最大的x的幂
			p := int(math.Log2(float64(curr)) / math.Log2(float64(x)))
			// 计算x^p的时候也可以考虑使用一个hash表，可以节省计算时间
			sum := exp(x, p)
			// x^p 所需要的运算符个数就是p个(0除外，p==0时，所需运算符的个数是2)，但是这里p不可能为0，只有当curr < x时p<0,但是已经进入上面的条件
			ans = p + find(curr-sum)
			if sum*x-curr < curr {
				ans = min(ans, p+1+find(sum*x-curr))
			}
		}
		table[curr] = ans
		return ans
	}
	return find(target) - 1
}

func exp(a int, e int) int {
	ret := 1
	for i := 1; i <= e; i++ {
		ret *= a
	}
	return ret
}
