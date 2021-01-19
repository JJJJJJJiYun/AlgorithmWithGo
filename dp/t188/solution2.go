package t188

// dp[i][0][j] = max(dp[i-1][0][j], dp[i-1][1][j+1]+price[i-1])
// dp[i][1][j] = max(dp[i-1][1][j], dp[i-1][0][j]-price[i-1])
func maxProfit2(k int, prices []int) int {
	var (
		notPossess = make([][]int, len(prices))
		possess    = make([][]int, len(prices))
	)
	for i := range notPossess {
		temp := make([]int, k+1)
		notPossess[i] = temp
	}
	for i := range possess {
		temp := make([]int, k+1)
		temp[0] = -9999
		possess[i] = temp
	}
	for i := 0; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			if i == 0 {
				notPossess[i][j] = 0
				possess[i][j] = -prices[i]
				continue
			}
			notPossess[i][j] = max(notPossess[i-1][j], possess[i-1][j]+prices[i])
			possess[i][j] = max(possess[i-1][j], notPossess[i-1][j-1]-prices[i])
		}
	}
	return notPossess[len(prices)-1][k]
}
