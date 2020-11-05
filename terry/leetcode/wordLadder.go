package leetcode

func LadderLength(beginWord string, endWord string, wordList []string) int {
	wordID := make(map[string]int)
	edges := [][]int{}
	wordNum := 0

	addWord := func(word string) int {
		if _, ok := wordID[word]; !ok {
			wordID[word] = wordNum
			edges = append(edges, []int{})
			wordNum++
		}
		return wordID[word]
	}
	addEdge := func(word string) {
		id1 := addWord(word)
		runes := []rune(word)
		for i := 0; i < len(runes); i++ {
			// 两个单词只有一个字母不同，那么这两个单词之间可以进行转换
			// 如果每次判断两个单词之间是否可以转换，那么会非常耗时
			// 我们对一个单词A(len(A)=L)创建L个虚拟节点，每个虚拟节点都是单词A替换一个字母为*的新单词
			// 那么只要两个单词有相同的虚拟节点，他们就可以进行转换
			tmp := runes[i]
			runes[i] = '*'
			id2 := addWord(string(runes))
			edges[id1] = append(edges[id1], id2)
			edges[id2] = append(edges[id2], id1)
			runes[i] = tmp
		}
	}
	for _, word := range wordList {
		addEdge(word)
	}
	addEdge(beginWord)
	if _, ok := wordID[endWord]; !ok {
		return 0
	}
	beginID, endID := wordID[beginWord], wordID[endWord]
	queue := []int{beginID}
	dist := make([]int, wordNum)
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			id := queue[i]
			if id == endID {
				// 由于有虚拟节点的存在，所以需要/2
				return dist[id]/2 + 1
			}
			for _, edge := range edges[id] {
				if dist[edge] == 0 {
					dist[edge] = dist[id] + 1
					queue = append(queue, edge)
				}
			}
		}
		queue = queue[l:]
	}
	return 0
}

// 从两端同时查找，可以有效降低广度优先搜索的每层节点分支
func LadderLengthI(beginWord string, endWord string, wordList []string) int {
	wordID := make(map[string]int)
	edges := [][]int{}
	wordNum := 0

	addWord := func(word string) int {
		if _, ok := wordID[word]; !ok {
			wordID[word] = wordNum
			edges = append(edges, []int{})
			wordNum++
		}
		return wordID[word]
	}
	addEdge := func(word string) {
		id1 := addWord(word)
		runes := []rune(word)
		for i := 0; i < len(runes); i++ {
			// 两个单词只有一个字母不同，那么这两个单词之间可以进行转换
			// 如果每次判断两个单词之间是否可以转换，那么会非常耗时
			// 我们对一个单词A(len(A)=L)创建L个虚拟节点，每个虚拟节点都是单词A替换一个字母为*的新单词
			// 那么只要两个单词有相同的虚拟节点，他们就可以进行转换
			tmp := runes[i]
			runes[i] = '*'
			id2 := addWord(string(runes))
			edges[id1] = append(edges[id1], id2)
			edges[id2] = append(edges[id2], id1)
			runes[i] = tmp
		}
	}
	for _, word := range wordList {
		addEdge(word)
	}
	if _, ok := wordID[endWord]; !ok {
		return 0
	}
	addEdge(beginWord)
	beginID, endID := wordID[beginWord], wordID[endWord]
	beginQueue := []int{beginID}
	endQueue := []int{endID}
	distBegin := make([]int, wordNum)
	distEnd := make([]int, wordNum)
	for len(beginQueue) > 0 && len(endQueue) > 0 {
		begin := beginQueue
		beginQueue = []int{}
		for _, id := range begin {
			if distEnd[id] != 0 {
				return (distEnd[id]+distBegin[id])/2 + 1
			}
			for _, edge := range edges[id] {
				if distBegin[edge] == 0 {
					distBegin[edge] = distBegin[id] + 1
					beginQueue = append(beginQueue, edge)
				}
			}
		}

		end := endQueue
		endQueue = []int{}
		for _, id := range end {
			if distBegin[id] != 0 {
				return (distBegin[id]+distEnd[id])/2 + 1
			}
			for _, edge := range edges[id] {
				if distEnd[edge] == 0 {
					distEnd[edge] = distEnd[id] + 1
					endQueue = append(endQueue, edge)
				}
			}
		}
	}
	return 0
}
