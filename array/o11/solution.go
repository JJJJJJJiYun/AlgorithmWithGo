package o11

// 旋转数组中的最小数字

// [1,2,3,4,5]
// [3,4,5,1,2]
// [1,2,3,4,5,6,7]
// [6,7,1,2,3,4,5]
func FindMinInRotatedArray(nums []int) int {
	i, j := 0, len(nums)-1
	for {
		mid := i + (j-i)/2
		if nums[mid] > nums[i] {
			i = mid
		} else {
			j = mid
		}
		if j-i == 1 {
			return nums[j]
		}
	}
}
