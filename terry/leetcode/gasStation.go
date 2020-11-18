package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	for i := 0; i < length; i++ {
		j := i
		left := gas[j]
		for left >= cost[j] {
			left = left - cost[j] + gas[(j+1)%length]
			j = (j + 1) % length
			if j == i {
				// 如果j==i那么说明已经绕了一圈
				return i
			}
		}
		if j < i {
			// j< i 说明绕到了i的前面，不可能有解
			return -1
		}
		j = i
	}
	return -1
}

func canCompleteCircuitI(gas []int, cost []int) int {
	// currentGas表示从开始处到现在车中的油剩余量
	// totalGas表示剩余油的总量
	totalGas, currentGas := 0, 0
	ret := 0
	for i := 0; i < len(gas); i++ {
		currentGas += gas[i] - cost[i]
		totalGas += gas[i] - cost[i]
		// 当currentGas < 0 时，说明无法通过，只能以下一个加油站为开始处
		if currentGas < 0 {
			ret = i + 1
			currentGas = 0
		}
	}
	// 只要 sum(gas) >= sum(cost) 就必然有解
	if totalGas < 0 {
		return -1
	}
	return ret
}
