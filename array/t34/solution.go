package t34

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left < len(nums) && nums[left] == target {
		res[0] = left
	}
	left, right = 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left > 0 && nums[left-1] == target {
		res[1] = left - 1
	}
	return res
}
