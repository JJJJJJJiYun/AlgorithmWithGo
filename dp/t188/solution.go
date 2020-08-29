package t188

import "math"

// 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k笔 交易。
// 注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

// dp[i][j][0] = max(dp[i-1][j][1]+price[i], dp[i-1][j][0])
// dp[i][j][1] = max(dp[i-1][j-1][0]-price[i], dp[i-1][j][1])
// dp[-1][j][0] = 0
// dp[-1][j][1] = -inf
// dp[i][0][0] = 0
// dp[i][0][1] = -inf
func maxProfit(k int, prices []int) int {
	// 注意这里，如果限制的次数超过了天数/2，说明其实并没有限制，因为至少买卖需要两天
	if k > len(prices)/2 {
		return maxProfitInf(prices)
	}
	dp := make([][][2]int, 0)
	for range prices {
		temp := make([][2]int, 0)
		for i := 0; i <= k; i++ {
			temp = append(temp, [2]int{})
		}
		dp = append(dp, temp)
	}
	for i := 0; i < len(prices); i++ {
		dp[i][0][0] = 0
		dp[i][0][1] = math.MinInt32
	}
	for i := 0; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}
			dp[i][j][0] = max(dp[i-1][j][1]+prices[i], dp[i-1][j][0])
			dp[i][j][1] = max(dp[i-1][j-1][0]-prices[i], dp[i-1][j][1])
		}
	}
	return dp[len(prices)-1][k][0]
}

func maxProfitInf(prices []int) int {
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
