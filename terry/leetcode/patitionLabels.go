package leetcode

func partitionLabels(S string) []int {
	length := len(S)
	// if length <= 1 {
	// 	return []int{1}
	// }
	parent := make([]int, length)
	union := make(map[rune]int)
	parent[0] = 1
	union[rune(S[0])] = 1
	for i := 1; i < length; i++ {
		tmp := rune(S[i])
		if union[tmp] > 0 {
			if union[tmp] == parent[i-1] {
				parent[i] = parent[i-1]
				continue
			}
			parent[i] = union[tmp]
			for j := 0; j < i; j++ {
				if parent[j] > union[tmp] {
					parent[j] = union[tmp]
					union[rune(S[j])] = union[tmp]
				}
			}

		} else {
			parent[i] = parent[i-1] + 1
			union[tmp] = parent[i]
		}
	}
	count := 0
	var prev int
	res := []int{}
	for i := 0; i < length; i++ {
		if prev == 0 {
			prev = parent[i]
			count++
		} else if prev == parent[i] {
			count++
		} else {
			res = append(res, count)
			count = 1
			prev = parent[i]
		}
	}
	return append(res, count)
}
