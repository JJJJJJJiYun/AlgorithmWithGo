package t18

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		threeSumRes := threeSum(nums[i+1:], target-num)
		for _, r := range threeSumRes {
			res = append(res, append(r, num))
		}
		for i+1 < len(nums) && nums[i+1] == num {
			i++
		}
	}
	return res
}

func threeSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		twoSumRes := twoSum(nums[i+1:], target-num)
		for _, r := range twoSumRes {
			res = append(res, append(r, num))
		}
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
