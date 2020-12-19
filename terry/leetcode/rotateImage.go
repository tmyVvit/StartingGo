package leetcode

// 将矩阵顺时针旋转90度
// 一层一层地计算
func rotate(matrix [][]int) {
	maxLength := len(matrix)
	if maxLength < 2 {
		return
	}
	start := 0
	// 每层的数组长度减2
	for count := maxLength - 1; count > 0; count -= 2 {
		rightlimit := count + start
		for i := 0; i < count; i++ {
			// tmp := matrix[start][start+i]
			// matrix[start][start+i] = matrix[rightlimit-i][start]
			// matrix[rightlimit-i][start] = matrix[rightlimit][rightlimit-i]
			// matrix[rightlimit][rightlimit-i] = matrix[start+i][rightlimit]
			// matrix[start+i][rightlimit] = tmp
			matrix[start][start+i], matrix[start+i][rightlimit], matrix[rightlimit][rightlimit-i], matrix[rightlimit-i][start] = matrix[rightlimit-i][start], matrix[start][start+i], matrix[start+i][rightlimit], matrix[rightlimit][rightlimit-i]
		}
		start++
	}
}

// n := length-1 -> 1 , n=n-2
//  i := 0 -> n-2
//  (0,0)     + (0,1)
//  (0,n-1)   + (1, 0)
//  (n-1,n-1) + (0, -1)
//  (n-1, 0)  + (-1, 0)
