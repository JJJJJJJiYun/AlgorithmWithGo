package t62

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
// 问总共有多少条不同的路径？

// 动态规划，dp[m][n]=dp[m-1][n]+dp[m][n-1]
// 更重要的是如何节约空间
// 使用一行数组，当前单元格就是dp[m-1][n],前一格就是dp[m][n-1]
func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i, _ := range dp {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}
