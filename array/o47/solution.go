package o47

// dp[m][n] = max(dp[m-1][n], dp[m][n-1])+gifts[m][n]
func MaxGiftValue(gifts [][]int) int {
	if gifts == nil || len(gifts) == 0 || gifts[0] == nil || len(gifts[0]) == 0 {
		return 0
	}
	dp := make([][]int, 0)
	for _, v := range gifts {
		temp := make([]int, 0)
		for range v {
			temp = append(temp, 0)
		}
		dp = append(dp, temp)
	}
	temp := 0
	for i, v := range gifts[0] {
		temp += v
		dp[0][i] = temp
	}
	temp = 0
	for i, v := range gifts {
		temp += v[0]
		dp[i][0] = temp
	}
	for i := 1; i < len(gifts); i++ {
		for j := 1; j < len(gifts[0]); j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + gifts[i][j]
		}
	}
	return dp[len(gifts)-1][len(gifts[0])-1]
}

func MaxGiftValue2(gifts [][]int) int {
	if gifts == nil || len(gifts) == 0 || gifts[0] == nil || len(gifts[0]) == 0 {
		return 0
	}
	dp := make([]int, len(gifts[0]))
	for i := 0; i < len(gifts); i++ {
		for j := 0; j < len(gifts[0]); j++ {
			up, left := 0, 0
			if i > 0 {
				up = dp[j]
			}
			if j > 0 {
				left = dp[j-1]
			}
			dp[j] = max(up, left) + gifts[i][j]
		}
	}
	return dp[len(gifts[0])-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
