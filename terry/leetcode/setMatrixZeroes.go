package leetcode

// 使用O(1)的额外空间
// 本方法使用两个额外的变量存储
// 首先如果matrix[i][j]==0，那么设置第i行和第j列的首位为0
// 对第0行和第0列特殊处理，分别记录在row0, col0
// 其实可以使用matrix[0][0]记录第0行，用col0记录第0列，这样可以再节省一个变量空间
func setZeros(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row0, col0 := false, false
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				if i == 0 && j == 0 {
					row0, col0 = true, true
				} else if i == 0 {
					row0 = true
				} else if j == 0 {
					col0 = true
				} else {
					matrix[i][0] = 0
					matrix[0][j] = 0
				}
			}
		}
	}
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
	if row0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if col0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
