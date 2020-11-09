package o40

// 最小的k个数

func MinK(nums []int, k int) []int {
	left, right := 0, len(nums)-1
	index := -1
	for {
		index = getIndex(nums, left, right)
		if index == k-1 {
			return nums[0:k]
		} else if index > k-1 {
			right = index - 1
		} else {
			left = index + 1
		}
	}
}

func getIndex(nums []int, left, right int) int {
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}
