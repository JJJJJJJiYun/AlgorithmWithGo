package microsoft

func sortColors(nums []int) {
	var i, j int
	for k := range nums {
		if nums[k] == 0 {
			nums[i], nums[k] = nums[k], nums[i]
			if i < j {
				nums[j], nums[k] = nums[k], nums[j]
			}
			i++
			j++
		} else if nums[k] == 1 {
			nums[j], nums[k] = nums[k], nums[j]
			j++
		}
	}
}
