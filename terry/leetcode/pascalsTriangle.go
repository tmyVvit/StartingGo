package leetcode

// 杨辉三角
func generate(numRows int) (ret [][]int) {
	for i := 0; i < numRows; i++ {
		currRow := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				currRow[j] = 1
			} else {
				currRow[j] = ret[i-1][j-1] + ret[i-1][j]
			}
		}
		ret = append(ret, currRow)
	}
	return ret
}
