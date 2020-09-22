package o1

// 在一个长度为n的数组中，元素都在0~n-1范围内，请找出其中的任意一个重复数字

// 改变原数组的方法
// 交换位置
func FindDuplicateNum(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}

// 不改变原数组的方法
// 二分查找
func FindDuplicateNum2(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + (j-i)/2
		count := countRangeNumInNums(nums, i, mid)
		if i == j && count > 1 {
			return i
		}
		if count > mid-i+1 {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return -1
}

// count数组中在[min,max]范围里数字的个数
func countRangeNumInNums(nums []int, min, max int) int {
	count := 0
	for _, num := range nums {
		if num >= min && num <= max {
			count++
		}
	}
	return count
}
