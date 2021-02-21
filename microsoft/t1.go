package microsoft

// 哈希，注意只需要一遍就可以
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		if index, ok := numsMap[target-num]; ok {
			return []int{i, index}
		}
		numsMap[num] = i
	}
	return nil
}
