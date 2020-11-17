package leetcode

import (
	"fmt"
)

// AllCellsDistOrder 一个R行 C列的矩阵，给定一个坐标（r0, c0），根据其他坐标到此坐标的距离返回所有其他坐标
// 使用 曼哈顿距离 |r0-r1| + |c0-c1|
// 使用BFS，广度优先
func AllCellsDistOrder(R int, C int, r0 int, c0 int) (res [][]int) {
	// 记录坐标是否已经被记录
	table := make(map[string]bool)
	queue := [][]int{[]int{r0, c0}}
	res = append(res, []int{r0, c0})
	table[getTalbeKey(r0, c0)] = true
	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			point := queue[i]
			for _, dir := range direction {
				r, c := point[0]+dir[0], point[1]+dir[1]
				key := getTalbeKey(r, c)
				if r >= 0 && r < R && c >= 0 && c < C && !table[key] {
					res = append(res, []int{r, c})
					queue = append(queue, []int{r, c})
					table[key] = true
				}
			}
		}
		queue = queue[length:]
	}
	return res
}

func getTalbeKey(r, c int) string {
	return fmt.Sprintf("%03d%03d", r, c)
}
