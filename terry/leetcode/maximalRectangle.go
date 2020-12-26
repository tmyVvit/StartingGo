package leetcode

// 暴力求解，遍历matrix，求以matrix[i][j]为长方形左上角顶点时的所有长方形的最大面积 -- （可以考虑以matrix[i][j]为长方形右下角顶点）
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	maxArea := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			maxArea = max(maxArea, findMaxAreaAt(matrix, i, j))
		}
	}
	return maxArea
}

func findMaxAreaAt(matrix [][]byte, x int, y int) int {
	if matrix[x][y] == 0 {
		return 0
	}
	var height, width, i, j, area int
	height = getHeight(matrix, x, y)
	width = getWidth(matrix, x, y)
	for i = 0; i < height; i++ {
		for j = 0; j < width; j++ {
			if isRec(matrix, x, y, i, j) {
				area = max(area, (1+i)*(1+j))
			}
		}
	}
	return area
}

func getHeight(matrix [][]byte, x int, y int) int {
	maxHeight := len(matrix)
	height := 0
	for x < maxHeight && matrix[x][y] == 1 {
		height++
		x++
	}
	return height
}

func getWidth(matrix [][]byte, x int, y int) int {
	maxWidth := len(matrix[x])
	width := 0
	for y < maxWidth && matrix[x][y] == 1 {
		width++
		y++
	}
	return width
}

func isRec(matrix [][]byte, x int, y int, i int, j int) bool {
	for m := 0; m <= i; m++ {
		for n := 0; n <= j; n++ {
			if matrix[x+m][y+n] == 0 {
				return false
			}
		}
	}
	return true
}
