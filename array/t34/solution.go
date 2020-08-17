package t34

// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
// 你的算法时间复杂度必须是O(log n) 级别。
// 如果数组中不存在目标值，返回[-1, -1]。

// 二分查找，注意循环结束条件，搜索范围变化，和最后的越界问题
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}
	left, right := 0, len(nums)
	// 二分查找最左边界
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left < len(nums) && nums[left] == target {
		res[0] = left
	}
	left, right = 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left > 0 && nums[left-1] == target {
		res[1] = left - 1
	}
	return res
}
