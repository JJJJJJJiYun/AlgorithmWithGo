package t122

import "math"

// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+price[i])
// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-price[i])
func maxProfit(prices []int) int {
	notPossess, possess := 0, math.MinInt32
	for i := 0; i < len(prices); i++ {
		lastNotPossess := notPossess
		notPossess = max(notPossess, possess+prices[i])
		possess = max(possess, lastNotPossess-prices[i])
	}
	return notPossess
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
