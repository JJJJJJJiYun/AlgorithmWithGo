package microsoft

func maxSubArray(nums []int) int {
	cur, res := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if cur+nums[i] > nums[i] {
			cur += nums[i]
		} else {
			cur = nums[i]
		}
		res = max(res, cur)
	}
	return res
}
