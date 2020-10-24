package t376

// 动态规划
// up[i]: nums[0:i]中以增大为结尾的最长摆动序列长度
// down[i]: nums[0:i]中以减小为结尾的最长摆动序列长度
// up[i] = down[i-1]+1 if nums[i]>nums[i-1]
// down[i] = up[i-1]+1 if nums[i]>nums[i-1]
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	up, down := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	if up > down {
		return up
	} else {
		return down
	}
}

// 贪心
// 实际上也就是找到波峰和波谷
func wiggleMaxLength2(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	count := 1
	if nums[1] != nums[0] {
		count = 2
	}
	prediff := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if (diff < 0 && prediff >= 0) || (diff > 0 && prediff <= 0) {
			count++
			prediff = diff
		}
	}
	return count
}
