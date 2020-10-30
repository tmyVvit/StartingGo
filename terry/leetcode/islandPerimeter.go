package leetcode

type pairs struct {
	x, y int
}

// 使用广度优先算法。 只要一个岛的一面是水，那么就可以算作周长
// 可以不用新建flag，直接在原本的grid做标记也是可以的
func islandPerimeter(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	flag := [][]int{}
	st := []pairs{}
	for i := 0; i < row; i++ {
		flagi := make([]int, col)
		flag = append(flag, flagi)
		for j := 0; j < col; j++ {
			flag[i][j] = grid[i][j]
			if grid[i][j] == 1 && len(st) == 0 {
				st = append(st, pairs{i, j})
			}
		}
	}
	perimeter := 0
	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(st) > 0 {
		land := st[0]
		st = st[1:]
		if flag[land.x][land.y] == 1 {
			flag[land.x][land.y] = 2
			for _, dir := range direction {
				x, y := land.x+dir[0], land.y+dir[1]
				if x < 0 || y < 0 || x >= row || y >= col || flag[x][y] == 0 {
					perimeter++
				} else if flag[x][y] == 1 {
					st = append(st, pairs{x, y})
				}
			}
		}
	}
	return perimeter
}

// 使用深度优先遍历DFS算法
func islandPerimeterI(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				return dfsToGetPerimeter(i, j, grid)
			}
		}
	}
	return 0
}

func dfsToGetPerimeter(x int, y int, grid [][]int) int {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
		return 1
	}
	if grid[x][y] == 2 {
		return 0
	}
	grid[x][y] = 2
	return dfsToGetPerimeter(x, y-1, grid) + dfsToGetPerimeter(x, y+1, grid) + dfsToGetPerimeter(x-1, y, grid) + dfsToGetPerimeter(x+1, y, grid)
}

// 直接遍历
func IslandPerimeterII(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	ret := 0
	for i := 0; i < row; i++ {
		for j := 0; j < row; j++ {
			if grid[i][j] == 1 {
				if i-1 < 0 || grid[i-1][j] == 0 {
					ret++
				}
				if i+1 >= row || grid[i+1][j] == 0 {
					ret++
				}
				if j-1 < 0 || grid[i][j-1] == 0 {
					ret++
				}
				if j+1 >= col || grid[i][j+1] == 0 {
					ret++
				}
			}
		}
	}
	return ret
}
