package t887

import "math"

// 你将获得 K 个鸡蛋，并可以使用一栋从 1 到 N  共有 N 层楼的建筑。
// 每个蛋的功能都是一样的，如果一个蛋碎了，你就不能再把它掉下去。
// 你知道存在楼层 F ，满足 0 <= F <= N 任何从高于 F 的楼层落下的鸡蛋都会碎，从 F 楼层或比它低的楼层落下的鸡蛋都不会破。
// 每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 X 扔下（满足 1 <= X <= N）。
// 你的目标是确切地知道 F 的值是多少。
// 无论 F 的初始值如何，你确定 F 的值的最小移动次数是多少？

// dp[i][j] = for k in [1,j] min(max(dp[i-1][k-1], dp[i][j-k])+1)
// dp[1][j] = j (0<=j<=N)
// dp[i][1] = 1 (1<=i<=K)
func superEggDrop(K int, N int) int {
	dp := make([][]int, 0)
	for i := 0; i <= K; i++ {
		temp := make([]int, 0)
		for j := 0; j <= N; j++ {
			temp = append(temp, math.MaxInt32)
		}
		dp = append(dp, temp)
	}
	for i := 0; i <= N; i++ {
		dp[1][i] = i
	}
	for i := 1; i <= K; i++ {
		dp[i][1] = 1
		dp[i][0] = 0
	}
	for i := 2; i <= K; i++ {
		for j := 2; j <= N; j++ {
			// 正常的dp
			//for k := 1; k <= j; k++ {
			//	 dp[i][j] = min(dp[i][j], max(dp[i-1][k-1], dp[i][j-k])+1)
			//}
			// 二分dp，真的很难理解
			left, right := 1, j
			for left < right {
				mid := left + (right-left+1)/2
				breakCount := dp[i-1][mid-1]
				notBreakCount := dp[i][j-mid]
				if breakCount > notBreakCount {
					right = mid - 1
				} else {
					left = mid
				}
			}
			dp[i][j] = max(dp[i-1][left-1], dp[i][j-left]) + 1
		}
	}
	return dp[K][N]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
