package leetcode

type union struct {
	nums []int
}

func (u *union) find(val int) int {
	if u.nums[val] == val {
		return u.nums[val]
	}
	return u.find(u.nums[val])
}

func (u *union) join(node1, node2 int) {
	parent1 := u.find(node1)
	parent2 := u.find(node2)
	if parent1 == parent2 {
		return
	}
	u.nums[parent1] = parent2
}

func (u *union) same(node1, node2 int) bool {
	return u.find(node1) == u.find(node2)
}

func initUnion(length int) *union {
	ancestors := make([]int, length)
	res := &union{ancestors}
	for i := 0; i < length; i++ {
		ancestors[i] = i
	}
	return res
}

func findRedundantDirectedConnection2(edges [][]int) []int {
	unions := initUnion(len(edges) + 1)
	parents := make([]int, len(edges)+1)
	for i := 0; i < len(edges)+1; i++ {
		parents[i] = i
	}
	conflict, cycle := -1, -1
	for idx, edge := range edges {
		u, v := edge[0], edge[1]
		if parents[v] != v { // 说明节点v已经有了父节点，出现冲突
			conflict = idx // 冲突的两条边分别是 [u, v] 和 [parents[v], v]
			continue       // 后出现的冲突不放入图中
		}
		parents[v] = u
		if unions.same(u, v) { // 如果u,v在相同的并差集中，说明出现了环，
			cycle = idx
		} else {
			unions.join(u, v) // 将u,v加入相同并差集中
		}
	}
	if conflict > -1 {
		if cycle < 0 {
			// 如果出现冲突，并且图中没有环，那么冲突的那一条边就是多余的
			return []int{edges[conflict][0], edges[conflict][1]}
		}
		// 如果出现冲突，并且图中还有环，那么出现环的那一条边就是多余的
		// 因为只多余一条，那么只会有两条边之间有冲突，因为有一条冲突的边并没有在图中，那么肯定是另一条导致的 : [parent[v], v]
		return []int{parents[edges[conflict][1]], edges[conflict][1]}
	}
	// 如果没有出现冲突，那么肯定有环出现，
	return []int{edges[cycle][0], edges[cycle][1]}
}
