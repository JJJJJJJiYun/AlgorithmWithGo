package sort

// BubbleSort 冒泡排序，每次遍历把最大的值交换到最后去
func BubbleSort(nums []int) {
	for i := range nums {
		isSwap := false
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				isSwap = true
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
		if !isSwap {
			break
		}
	}
}
