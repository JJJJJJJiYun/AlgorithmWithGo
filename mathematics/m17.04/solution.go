package m17_04

// 异或！
func missingNumber(nums []int) int {
	//sum, all := 0, 0
	//for i, num := range nums {
	//	sum += num
	//	all += i + 1
	//}
	//return all-sum
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= i
		res ^= nums[i]
	}
	res ^= len(nums)
	return res
}
