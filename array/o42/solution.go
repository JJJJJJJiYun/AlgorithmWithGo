package o42

import (
	"math"
)

// 连续子数组的最大和

func MaxSubArraySum(nums []int) int {
	max := math.MinInt32
	cur := math.MinInt32
	for _, num := range nums {
		if cur < 0 && num > cur {
			cur = num
		} else if cur >= 0 {
			cur += num
		}
		if cur > max {
			max = cur
		}
	}
	return max
}

// f(i)表示以num[i]结尾的最大值
// f(i) = f(i-1) + nums[i] if f(i-1)>=0
// f(i) = nums[i] if f(i-1)<0
// 结果是max(f(i))
