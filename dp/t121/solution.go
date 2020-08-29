package t121

import "math"

// 股票问题题解
// 利用状态递归，状态有：天数、可交易次数、持有状态（买和卖是动作，不是状态）
// dp = [天数][可交易次数][是否持有]int{当前最大利润}
// dp[i][j][0] = max(dp[i-1][j][1]+price[i], dp[i-1][j][0])
// dp[i][j][1] = max(dp[i-1][j-1][0]-price[i], dp[i-1][j][1])
// dp[-1][j][0] = 0
// dp[-1][j][1] = -inf
// dp[i][0][0] = 0
// dp[i][0][1] = -inf

// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
// 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
// 注意：你不能在买入股票前卖出股票。

// dp[i][0] = max(dp[i-1][1]+price[i], dp[i-1][0])
// dp[i][1] = max(-price[i], dp[i-1][1])
func maxProfit(prices []int) int {
	notPossess, possess := 0, math.MinInt32
	for i := 0; i < len(prices); i++ {
		notPossess = max(possess+prices[i], notPossess)
		possess = max(-prices[i], possess)
	}
	return notPossess
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
