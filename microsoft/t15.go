package microsoft

import "sort"

func threeSum(nums []int) (res [][]int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		k := len(nums) - 1
		for j := i + 1; j < len(nums)-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for ; k > j && nums[j]+nums[k] > target; k-- {
			}
			if k == j {
				break
			}
			if nums[j]+nums[k] == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return res
}
