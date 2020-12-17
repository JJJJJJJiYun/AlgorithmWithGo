package t714

// dp[i][0] = max(dp[i-1][0],dp[i-1][1]+price[i]-fee)
// dp[i][1] = max(dp[i-1][1],dp[i-1][0]-price[i])
func maxProfit(prices []int, fee int) int {
	possess, notPossess := -99999, 0
	for _, p := range prices {
		lastPossess := possess
		possess = max(possess, notPossess-p)
		notPossess = max(notPossess, lastPossess+p-fee)
	}
	return notPossess
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
