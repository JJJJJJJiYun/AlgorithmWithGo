package t198

// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
// 影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

// dp[i] = max(dp[i+1], dp[i+2]+nums[i])
func rob(nums []int) int {
	next, nnext := 0, 0
	res := 0
	for i := len(nums) - 1; i >= 0; i-- {
		res = max(next, nnext+nums[i])
		nnext = next
		next = res
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
