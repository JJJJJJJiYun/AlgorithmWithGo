package o14

import "LeetCode/utils"

// dp[n] = max(dp[n-i]*dp[i])
func CutRope(n int) int {
	if n <= 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	dp := make([]int, 0)
	for i := 0; i <= n; i++ {
		dp = append(dp, 0)
	}
	// 这个地方的初始值，实际上不是长度为n时的题解
	// 因为题目要求是必须剪一刀，但这个时候默认是可以不剪的
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = utils.Max(dp[i], dp[i-j]*dp[j])
		}
	}
	return dp[n]
}
