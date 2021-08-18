package t215

func findKthLargest(nums []int, k int) int {
	return helper(nums, 0, len(nums)-1, k)
}

func helper(nums []int, left, right, k int) int {
	i := getIndex(nums, left, right)
	if i == k-1 {
		return nums[i]
	} else if i > k-1 {
		return helper(nums, left, i-1, k)
	}
	return helper(nums, i+1, right, k)

}

func getIndex(nums []int, left, right int) int {
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] <= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] >= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}
