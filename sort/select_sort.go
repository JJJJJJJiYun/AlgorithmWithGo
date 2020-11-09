package sort

// SelectSort 选择排序，每次找到未排序部分的最小值放到当前位置
func SelectSort(nums []int) {
	for i := range nums {
		minIndex := i
		for j := i; j < len(nums); j++ {
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}
