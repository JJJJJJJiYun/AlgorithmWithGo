package t1024

import "math"

// [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]]
// 10
// dp[i]表示能够覆盖[0,i]的最少个数
// dp[i] = min(dp[j])+1 if exist [j,k] && k>=i
func videoStitching(clips [][]int, T int) int {
	dp := make([]int, T+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for _, pair := range clips {
			if pair[0] < i && pair[1] >= i && dp[pair[0]]+1 < dp[i] {
				dp[i] = dp[pair[0]] + 1
			}
		}
	}
	if dp[T] == math.MaxInt32 {
		return -1
	}
	return dp[T]
}

func videoStitching2(clips [][]int, T int) int {
	maxn := make([]int, T)
	// 找到每个从i开始的能覆盖到的最远索引
	for _, pair := range clips {
		if pair[0] < T && maxn[pair[0]] < pair[1] {
			maxn[pair[0]] = pair[1]
		}
	}
	last, pre, count := 0, 0, 0
	for i, v := range maxn {
		if v > last {
			last = v
		}
		// 如果已经走到了最远索引，说明达不到了
		if i == last {
			return -1
		}
		// 如果走到了上一个区间的最大索引，说明要用下一个区间了
		if i == pre {
			count++
			pre = last
		}
	}
	return count
}
