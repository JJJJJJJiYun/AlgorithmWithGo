package microsoft

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i, num := range nums {
		if num <= 0 {
			nums[i] = n + 1
		}
	}
	for i := range nums {
		index := abs(nums[i])
		if index > 0 && index < n+1 && nums[index-1] > 0 {
			nums[index-1] = -nums[index-1]
		}
	}
	for i, num := range nums {
		if num > 0 {
			return i + 1
		}
	}
	return n + 1
}
