package t674

func findLengthOfLCIS(nums []int) int {
	// 如果长度小于等于1，最长为长度
	if len(nums) <= 1 {
		return len(nums)
	}
	i, j := 0, 1
	res := -1
	for ; j < len(nums); j++ {
		if nums[j] > nums[j-1] {
			if res < j-i {
				res = j - i
			}
		} else {
			i = j
		}
	}
	if res == -1 {
		return 1
	}
	return res
}
