package leetcode

import "sort"

type peoples struct {
	people [][]int
}

func (p *peoples) Less(a, b int) bool {
	if p.people[a][0] == p.people[b][0] {
		return p.people[a][1] > p.people[b][1]
	}
	return p.people[a][0] < p.people[b][0]
}
func (p *peoples) Swap(a, b int) { p.people[a], p.people[b] = p.people[b], p.people[a] }
func (p *peoples) Len() int      { return len(p.people) }

// ReconstructQueue 对队列中的人重新排序 【height, number】 height为身高，number是前面的人中更高或者身高相同的人的数量
// 先按身高升序，k值降序排列，在对第i个人排列时0,..i-1个人不会造成影响，我们只需要在前面留k_i个空位就好
func ReconstructQueue(people [][]int) [][]int {
	flag := make(map[int]bool)

	ps := &peoples{people}
	sort.Sort(ps)
	res := make([][]int, len(people))
	for _, p := range ps.people {
		var empty, i int
		for ; ; i++ {
			// 第i位还没有人时 或者 有人但是身高相等时 （因为前面已经按照身高排序，所以不可能有身高更高的人）
			if !flag[i] {
				if empty == p[1] {
					break
				} else {
					empty++
				}
			}

		}
		res[i] = p
		flag[i] = true
	}
	return res
}

// ReconstructQueueI 先按身高降序，k值升序排列。在排列第i个人时，我们只需要将这个人插入第i个位置即可
func ReconstructQueueI(people [][]int) (res [][]int) {
	sort.Slice(people, func(a, b int) bool {
		if people[a][0] == people[b][0] {
			return people[a][1] < people[b][1]
		}
		return people[a][0] > people[b][0]
	})
	for _, person := range people {
		idx := person[1]
		res = append(res[:idx], append([][]int{person}, res[idx:]...)...)
	}
	return res
}
