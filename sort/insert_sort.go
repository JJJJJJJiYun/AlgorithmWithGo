package sort

// InsertSort 插入排序，每次把当前位置的值插入到前面的有序部分中
func InsertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > num {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = num
	}
}
