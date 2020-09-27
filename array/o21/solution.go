package o21

// 调整顺序使奇数位于偶数前面

func ReorderOddAndEven(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[i]%2 == 1 {
			i++
		}
		for i < j && nums[j]%2 == 0 {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}
