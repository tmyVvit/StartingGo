package leetcode

import "math"

// CrackSafe n代表密码位数，k代表每位密码的值是0,1,...,k-1 (可以看作是k进制)
func CrackSafe(n int, k int) string {
	// 我们将所有n-1位数当作节点，那么可以有k^(n-1)个节点
	// 每个节点有k条边，一个节点a1a2..an-1的第x条边连接的就是a2a3...an-1x这个节点 （类似左移操作，并且加上边的值）
	// 我们可以把每个节点代表的值当作是一个k进制的值
	kn, kn1 := pow(k, n), pow(k, n-1)
	num := make([]int, kn1)
	// 最终结果肯定是 k^n + n-1 位
	ans := make([]byte, kn+n-1)
	for i := 0; i < kn+n-1; i++ {
		if i < kn1 {
			// 每个节点的可选择的边可以从k-1 -> 0，我们从大往小选择
			num[i] = k - 1
		}
		ans[i] = '0'
	}
	node := 0
	// 从第0个节点开始
	for i := n - 1; i < len(ans); i++ {
		ans[i] = byte('0' + num[node])
		num[node]--
		// 计算连接到的节点，左移并加上边的值
		node = (node*k + num[node] + 1) % kn1
	}
	return string(ans)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
