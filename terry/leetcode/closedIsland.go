package leetcode

import "fmt"

func main() {
	grid := [][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}}
	//grid := [][]int {{1,1,1},{1,0,1},{1,1,1}}
	fmt.Println(closedIsland(grid))
}

func closedIsland(grid [][]int) (ret int) {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return ret
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				ret += dfs(grid, i, j)
				printGrid(grid)
			}
		}
	}
	return ret
}

func dfs(grid [][]int, i int, j int) (ret int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return 0
	}
	if grid[i][j] > 0 {
		return 1
	}
	ret = 1
	grid[i][j] = 2
	x := []int{0, 0, 1, -1}
	y := []int{1, -1, 0, 0}
	for k := 0; k < 4; k++ {
		ret = min(ret, dfs(grid, i+x[k], j+y[k]))
	}
	return ret
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func printGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
