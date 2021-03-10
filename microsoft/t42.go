package microsoft

func trap(height []int) int {
	sum := 0
	maxLeft, maxRight := make([]int, len(height)), make([]int, len(height))
	for i := 1; i < len(height); i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i+1])
	}
	for i := 1; i < len(height)-1; i++ {
		min := min(maxLeft[i], maxRight[i])
		if min > height[i] {
			sum += min - height[i]
		}
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
