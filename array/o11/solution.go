package o11

import "math"

// 旋转数组中的最小数字

// [1,2,3,4,5]
// [3,4,5,1,2]
// [1,2,3,4,5,6,7]
// [6,7,1,2,3,4,5]
func FindMinInRotatedArray(nums []int) int {
	i, j := 0, len(nums)-1
	// 说明是顺序排列的
	if nums[i] < nums[j] {
		return nums[i]
	}
	// 是旋转过的
	for {
		mid := i + (j-i)/2
		if j-i == 1 {
			return nums[j]
		}
		if nums[i] == nums[j] && nums[i] == nums[mid] {
			min := math.MaxInt32
			for _, num := range nums {
				if num < min {
					min = num
				}
			}
			return min
		}
		if nums[mid] >= nums[i] {
			i = mid
		} else if nums[mid] <= nums[j] {
			j = mid
		}
	}
}
