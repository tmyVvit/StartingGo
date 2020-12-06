package leetcode

func destCity(paths [][]string) string {
	cityRoute := make(map[string]string)
	for _, path := range paths {
		cityRoute[path[0]] = path[1]
	}
	start := paths[0][0]
	for {
		if dest, ok := cityRoute[start]; ok {
			start = dest
		} else {
			break
		}
	}
	return start
}
