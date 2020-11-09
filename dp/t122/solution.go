package t122

import "math"

// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

func maxProfit(prices []int) int {
	// dp[i][true] = max(dp[i-1][true],dp[i-1][false])-nums[i]）
	// dp[i][false] = max(dp[i-1][true]+nums[i], dp[i-1][false])
	possess, notPossess := math.MinInt32, 0
	for _, num := range prices {
		lastPossess := possess
		possess = max(possess, notPossess-num)
		notPossess = max(lastPossess+num, notPossess)
	}
	return notPossess
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
