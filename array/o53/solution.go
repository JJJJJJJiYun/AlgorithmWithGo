package o53

func CountNumInSortedArray(nums []int, num int) int {
	return findRightBorder(nums, num) - findLeftBorder(nums, num) + 1
}

// 1,1,2,3,3,3,4,5,5
func findLeftBorder(nums []int, num int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] < num {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	if i >= 0 && nums[i] == num {
		return i
	}
	return -1
}

func findRightBorder(nums []int, num int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] <= num {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	if i-1 <= len(nums)-1 && nums[i-1] == num {
		return i - 1
	}
	return -1
}

// 一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。
// 在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。

// 二分查找
func missingNumber(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := (i + j) / 2
		if m == nums[m] {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return i
}

// 我的最初解法，递归二分，其实思路很像，但不够简洁，需要总结二分查找的解题规律
func missingNumber2(nums []int) int {
	return find(nums, 0, len(nums))
}

func find(nums []int, min, max int) int {
	target := (min + max) / 2
	if len(nums) == 1 {
		if nums[0] <= target {
			return nums[0] + 1
		} else {
			return nums[0] - 1
		}
	}
	if nums[len(nums)/2] <= target {
		return find(nums[len(nums)/2:], target, max)
	} else {
		return find(nums[0:len(nums)/2], min, target)
	}
}
