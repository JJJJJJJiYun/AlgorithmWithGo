package t31

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i := len(nums) - 1
	for i >= 1 && nums[i] <= nums[i-1] {
		i--
	}
	if i == len(nums)-1 {
		nums[i], nums[i-1] = nums[i-1], nums[i]
		return
	}
	if i == 0 {
		j := len(nums) - 1
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
		return
	}
	j := i
	for ; j < len(nums) && nums[j] > nums[i-1]; j++ {
	}
	j--
	nums[j], nums[i-1] = nums[i-1], nums[j]
	j = len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
