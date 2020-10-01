package o39

func MoreThanHalfNum(nums []int) int {
	i, j := 0, len(nums)-1
	index := -1
	for {
		index = getIndex(nums, i, j)
		if index < len(nums)/2 {
			i = index + 1
		} else if index > len(nums)/2 {
			j = index - 1
		} else {
			break
		}
	}
	return nums[index]
}

func MoreThanHalfNum2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		panic("invalid input")
	}
	count := 1
	num := nums[0]
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			num = nums[i]
			count++
		} else if num == nums[i] {
			count++
		} else {
			count--
		}
	}
	return num
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
