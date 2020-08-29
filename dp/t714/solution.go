package t714

import "math"

// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+price[i])
// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-price[i]-fee)
func maxProfit(prices []int, fee int) int {
	notPossess, possess := 0, math.MinInt32
	for i := 0; i < len(prices); i++ {
		lastNotPossess := notPossess
		notPossess = max(notPossess, possess+prices[i])
		possess = max(possess, lastNotPossess-prices[i]-fee)
	}
	return notPossess
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
