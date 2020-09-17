package main

import "fmt"

// ugly solution !!
func main() {
	graph := [][]int{{4, 2}, {1, 5}, {5, 2}, {5, 3}, {2, 4}}
	fmt.Println(findRedundantDirectedConnection(graph))
}

func findRedundantDirectedConnection(edges [][]int) (res []int) {
	inNode := make([]int, len(edges))
	outNode := make([]int, len(edges))
	grid := make([][]int, len(edges))
	for i := 0; i < len(edges); i++ {
		grid[i] = make([]int, len(edges))
	}
	for _, value := range edges {
		inNode[value[1]-1]++
		outNode[value[0]-1]++
		grid[value[0]-1][value[1]-1] = 1
	}

	fmt.Printf(" in node: %v\n", inNode)
	fmt.Printf("out node: %v\n", outNode)
	for i := len(edges) - 1; i >= 0; i-- {
		fmt.Println("-----------")
		fmt.Printf("  remove: %v\n", edges[i])
		grid[edges[i][0]-1][edges[i][1]-1] = 0
		inNode[edges[i][1]-1]--
		outNode[edges[i][0]-1]--
		fmt.Printf(" in node: %v\n", inNode)
		fmt.Printf("out node: %v\n", outNode)
		if isGraphRootedTree(inNode, outNode, grid) {
			return edges[i]
		}
		inNode[edges[i][1]-1]++
		outNode[edges[i][0]-1]++
		grid[edges[i][0]-1][edges[i][1]-1] = 1
	}
	return res
}

func isGraphRootedTree(inNode []int, outNode []int, grid [][]int) bool {
	var rootNum int
	var root int
	for i := 0; i < len(inNode); i++ {
		if (inNode[i] == 0 && outNode[i] == 0) || inNode[i] > 1 {
			return false
		}
		if inNode[i] == 0 {
			rootNum++
			root = i
		}
		if rootNum >= 2 {
			return false
		}
	}
	return isThroughGraph(grid, root)
}

func isThroughGraph(grid [][]int, root int) bool {
	st := new(stack)
	mask := make([]int, len(grid))
	mask[root] = 1
	for i := 0; i < len(grid[root]); i++ {
		if grid[root][i] == 1 && mask[i] == 0 {
			mask[i] = 1
			st.push(i)
		}
	}
	parentNodes := make([]int, len(grid))
	for k := st.pop(); k >= 0; k = st.pop() {
		if parentNodes[k] == 1 {
			continue
		}
		parentNodes[k] = 1
		for i := 0; i < len(grid[k]); i++ {
			if grid[k][i] == 1 && mask[i] == 0 {
				mask[i] = 1
				if parentNodes[i] == 0 {
					st.push(i)
				}
			}
		}
	}
	for i := 0; i < len(mask); i++ {
		if mask[i] == 0 {
			return false
		}
	}
	return true
}

type stack struct {
	index int
	list  []int
}

func (st *stack) pop() int {
	if st.index == 0 {
		return -1
	}
	st.index--
	return st.list[st.index]
}

func (st *stack) push(val int) {
	if len(st.list) > st.index {
		st.list[st.index] = val
	} else {
		st.list = append(st.list, val)
	}
	st.index++
}
