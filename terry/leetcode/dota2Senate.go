package leetcode

func PredictPartyVictory(senate string) string {
	dire, radiant := []int{}, []int{}
	length := len(senate)
	for idx, ch := range senate {
		if ch == 'D' {
			dire = append(dire, idx)
		} else if ch == 'R' {
			radiant = append(radiant, idx)
		}
	}
	// 每个参议员都会优先禁止下一次最先投票的对方议员
	// 如果此轮后面没有对方议员的话，就会禁止下一轮中最先投票的对方议员
	for len(dire) > 0 && len(radiant) > 0 {
		if dire[0] < radiant[0] {
			// 被禁止的议员永久出列
			radiant = radiant[1:]
			// 当前议员出列后，增加 length 然后入队列，确保议员的投票顺序
			tmp := dire[0] + length
			dire = append(dire[1:], tmp)
		} else {
			dire = dire[1:]
			tmp := radiant[0] + length
			radiant = append(radiant[1:], tmp)
		}
	}
	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}
