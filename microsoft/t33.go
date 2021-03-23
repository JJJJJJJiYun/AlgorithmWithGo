package microsoft

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		// 判断是否在前半部分查找，这里枚举出所有可能性
		if (nums[left] <= target && target <= nums[mid]) ||
			(nums[mid] <= nums[right] && (target < nums[mid] ||
				target > nums[right])) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
