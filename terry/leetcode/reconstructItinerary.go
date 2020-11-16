package leetcode

import "sort"

// 重新安排行程
func findItinerary(tickets [][]string) (res []string) {
	itineraries := make(map[string][]string)

	for _, tickets := range tickets {
		itineraries[tickets[0]] = append(itineraries[tickets[0]], tickets[1])
	}
	for _, itinerary := range itineraries {
		// 按照字符排序
		sort.Strings(itinerary)
	}
	var dfs func(from string)
	dfs = func(from string) {
		for {
			if itinerary, ok := itineraries[from]; !ok || len(itinerary) == 0 {
				break
			}
			to := itineraries[from][0]
			itineraries[from] = itineraries[from][1:]
			dfs(to)
		}
		// for循环结束后才入栈，则最后返回的先入栈
		res = append(res, from)
	}
	dfs("JKS")
	length := len(res)
	for i := 0; i <= (length-1)/2; i++ {
		res[i], res[length-1-i] = res[length-1-i], res[i]
	}
	return res
}
