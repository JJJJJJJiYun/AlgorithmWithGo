package t922

func sortArrayByParityII(nums []int) []int {
	i, j := 0, 1
	for {
		for i < len(nums) {
			if nums[i]%2 != 0 {
				break
			}
			i += 2
		}
		if i >= len(nums) {
			return nums
		}
		for j < len(nums) {
			if nums[j]%2 == 0 {
				break
			}
			j += 2
		}
		if j >= len(nums) {
			return nums
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
}
