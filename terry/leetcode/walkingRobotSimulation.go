package leetcode

import "fmt"

func robotSim(commands []int, obstacles [][]int) int {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dir := 0
	destX, destY := 0, 0
	res := 0
	for _, command := range commands {
		if command == -1 {
			dir = (dir + 1) % 4
		} else if command == -2 {
			dir = (dir + 3) % 4
		} else {
			tmpX, tmpY := destX+directions[dir][0]*command, destY+directions[dir][1]*command
			destX, destY = goObstables(obstacles, destX, destY, tmpX, tmpY)
			res = max(res, destX*destX+destY*destY)
		}
	}
	return res
}

func goObstables(obstacles [][]int, startx int, starty int, endx int, endy int) (int, int) {
	if len(obstacles) == 0 {
		return endx, endy
	} else if startx == endx {
		for _, obstacle := range obstacles {
			if obstacle[0] == startx {
				if (obstacle[1] > starty && obstacle[1] <= endy) || (obstacle[1] < starty && obstacle[1] >= endy) {
					if starty > endy {
						return startx, obstacle[1] + 1
					}
					if starty < endy {
						return startx, obstacle[1] - 1
					}
				}

			}
		}
	} else if starty == endy {
		for _, obstacle := range obstacles {
			if obstacle[1] == starty {
				if (obstacle[0] > startx && obstacle[0] <= endx) || (obstacle[0] < startx && obstacle[0] >= endx) {
					if startx > endx {
						return obstacle[0] + 1, endy
					}
					if startx < endx {
						return obstacle[0] - 1, endy
					}
				}
			}
		}
	}
	return endx, endy
}

// 每次执行一步，判断是否有陷阱
// 将陷阱数组编码，并存储进map中，可以更快
func robotSimI(commands []int, obstacles [][]int) int {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dir := 0
	destX, destY := 0, 0
	// 可以考虑使用map[int32]int类型的map，速度更快 使用encrypt1()函数
	obsMaps := make(map[string]int)
	for _, obstacle := range obstacles {
		tmp := encrypt(obstacle)
		obsMaps[tmp] = 1
	}
	res := 0
	for _, command := range commands {
		if command == -1 {
			dir = (dir + 1) % 4
		} else if command == -2 {
			dir = (dir + 3) % 4
		} else {
			for i := 0; i < command; i++ {
				tmpx, tmpy := destX+directions[dir][0], destY+directions[dir][1]
				if obsMaps[encrypt([]int{tmpx, tmpy})] == 1 {
					break
				}
				destX, destY = tmpx, tmpy
			}
			res = max(res, destX*destX+destY*destY)
		}
	}
	return res
}

func encrypt(x []int) string {
	return fmt.Sprintf("%06d", x[0]) + fmt.Sprintf("%06d", x[1])
}

func encrypt1(x []int) int32 {
	return int32(x[0])<<16 + int32(x[1])
}
