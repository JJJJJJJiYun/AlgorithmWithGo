package t15

import (
	"sort"
)

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。

// 利用twoSum，但要注意去重
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		// 只在比这个数大的数字中去找
		twoSumRes := twoSum(nums[i+1:], -num)
		for _, r := range twoSumRes {
			res = append(res, append(r, num))
		}
		// 重复数字跳过
		for i+1 < len(nums) && nums[i+1] == num {
			i++
		}
	}
	return res
}

func twoSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	i, j := 0, len(nums)-1
	sort.Ints(nums)
	for i < j {
		left, right := nums[i], nums[j]
		sum := left + right
		if sum < target {
			for nums[i] == left && i < j {
				i++
			}
		} else if sum > target {
			for nums[j] == right && i < j {
				j--
			}
		} else {
			res = append(res, []int{left, right})
			for nums[i] == left && i < j {
				i++
			}
			for nums[j] == right && i < j {
				j--
			}
		}
	}
	return res
}
