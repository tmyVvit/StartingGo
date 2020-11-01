package leetcode

// 使用动态规划 dp[i]代表s的前i个可以合法拆分
// 计算dp[i]时需要遍历 0 <= j < i，dp[j]==true并且s[j:i]是合法的一个单词
func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
