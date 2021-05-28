package microsoft

// dp[i] 表示s的前i个字符是否能够用wordDict中的单词组成
// dp[i] = for {dp[j] && check(s[j:i+1])}
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, w := range wordDict {
		wordMap[w] = true
	}
	wordMap[""] = true
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if dp[j] {
				if !(j <= i && !wordMap[s[j:i+1]]) {
					dp[i+1] = true
				}
			}
		}
	}
	return dp[len(s)]
}
