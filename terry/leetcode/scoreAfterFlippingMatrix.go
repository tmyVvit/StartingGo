package leetcode

// 给定一个翻转方案，则它们之间任意交换顺序后，得到的结果保持不变。因此，我们总可以先考虑所有的行翻转，再考虑所有的列翻转
// 为了得到最高的分数，矩阵的每一行的最左边的数都必须为 1
// 然后就需要进行列翻转，为了使总分最大，我们需要使每一列中的1数量最多
func matrixScore(A [][]int) int {
	if len(A) == 0 {
		return 0
	}
	rows, cols := len(A), len(A[0])
	// 我们不需要真正改变原数组，只需要计算出翻转之后的值即可
	// 首先将第一列全部置为1
	ans := rows * (1 << (cols - 1))
	// 然后对其余列进行判断，如果0的个数大于1，则翻转该列，否则不翻转
	for j := 1; j < cols; j++ {
		currColOne := 0
		for i := 0; i < rows; i++ {
			// 如果第一列的值是1，则代表没有翻转
			if A[i][0] == 1 {
				currColOne += A[i][j]
			} else {
				// 第一列的值是0，已经翻转了，那么此时的值就是 1 - A[i][j]
				currColOne += (1 - A[i][j])
			}
		}
		// currColOne 为该列 1 的个数，
		ans += max(currColOne, rows-currColOne) * (1 << (cols - j - 1))
	}
	return ans
}
