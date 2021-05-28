package microsoft

// [1,2,3,4]
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(ans); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	r := 1
	for i := len(ans) - 1; i >= 0; i-- {
		ans[i] = ans[i] * r
		r *= nums[i]
	}
	return ans
}
