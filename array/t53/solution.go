package t53

import "math"

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	cur := -1
	max := math.MinInt32
	for _, num := range nums {
		if cur < 0 {
			cur = num
		} else {
			cur += num
		}
		if max < cur {
			max = cur
		}
	}
	return max
}
