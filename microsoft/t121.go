package microsoft

// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+price[i])
// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-price[i])
func maxProfit(prices []int) int {
	p, np := -10000, 0
	for _, price := range prices {
		np = max(np, p+price)
		p = max(p, -price) // 如果只允许买卖一次，那么第一次持有一定是-price
	}
	return np
}
