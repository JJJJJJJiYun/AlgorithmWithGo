package microsoft

func maxSubArray(nums []int) int {
	curMax := 0
	max := nums[0]
	for _, num := range nums {
		if curMax < 0 && num > curMax {
			curMax = num
		} else {
			curMax += num
		}
		if max < curMax {
			max = curMax
		}
	}
	return max
}
