package t309

import "math"

// 给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
// 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+price[i])
// dp[i][1] = max(dp[i-1][1], dp[i-2][0]-price[i])
func maxProfit(prices []int) int {
	lastNotPossess, notPossess, possess := 0, 0, math.MinInt32
	for i := 0; i < len(prices); i++ {
		// 记录昨天
		temp := notPossess
		// 更新今天
		notPossess = max(notPossess, possess+prices[i])
		possess = max(possess, lastNotPossess-prices[i])
		// 记录前天
		lastNotPossess = temp
	}
	return notPossess
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
