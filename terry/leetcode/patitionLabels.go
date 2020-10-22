package leetcode

// PartitionLabels 使用并差集的思想
// 没访问一个字母，就会判断该字母所属的并差集
// 如果是新出现的字母，则新增一个并差集
// 如果是已存在的，并且属于最后一个并差集，则继续
// 如果是已存在的，并且不属于最后一个并差集，则将该并差集和最后一个并差集之间的所有元素合并。
func PartitionLabels(S string) []int {
	length := len(S)
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

// PartitionLabelsI 也是使用并差集思想
func PartitionLabelsI(S string) (res []int) {
	union := make(map[string]int)
	reversUnion := make(map[int]string)
	parts := 0
	for i := 0; i < len(S); i++ {
		ch := string(S[i])
		if union[ch] == 0 {
			parts++
			union[ch] = parts
			reversUnion[parts] = ch
		} else if union[ch] == parts {
			reversUnion[parts] = reversUnion[parts] + ch
		} else {
			for j := union[ch] + 1; j <= parts; j++ {
				reversUnion[union[ch]] = reversUnion[union[ch]] + reversUnion[j]
				delete(reversUnion, j)
			}
			for key, value := range union {
				if value > union[ch] {
					union[key] = union[ch]
				}
			}
			parts = union[ch]
			reversUnion[parts] = reversUnion[parts] + ch
		}
	}
	for i := 1; i <= parts; i++ {
		res = append(res, len(reversUnion[i]))
	}
	return res
}
