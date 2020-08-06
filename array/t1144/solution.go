package t1144

// 给你一个整数数组 nums，每次 操作 会从中选择一个元素并 将该元素的值减少 1。
// 如果符合下列情况之一，则数组 A 就是 锯齿数组：
// 每个偶数索引对应的元素都大于相邻的元素，即 A[0] > A[1] < A[2] > A[3] < A[4] > ...
// 或者，每个奇数索引对应的元素都大于相邻的元素，即 A[0] < A[1] > A[2] < A[3] > A[4] < ...
// 返回将数组 nums 转换为锯齿数组所需的最小操作次数。

// 分奇偶讨论就可以了
func movesToMakeZigzag(nums []int) int {
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	step1, step2 := 0, 0
	for i, num := range nums {
		var m int
		if i == 0 {
			m = min(1001, nums[i+1])
		} else if i == len(nums)-1 {
			m = min(nums[i-1], 1001)
		} else {
			m = min(nums[i-1], nums[i+1])
		}
		// 奇数大的,偶数才减
		if i%2 == 1 && num >= m {
			step1 += num - m + 1
			nums[i] = m - 1
		}
	}
	for i, num := range copyNums {
		var m int
		if i == 0 {
			m = min(1001, copyNums[i+1])
		} else if i == len(copyNums)-1 {
			m = min(copyNums[i-1], 1001)
		} else {
			m = min(copyNums[i-1], copyNums[i+1])
		}
		// 偶数大的,奇数才减
		if i%2 == 0 && num >= m {
			step2 += num - m + 1
			copyNums[i] = m - 1
		}
	}

	return min(step1, step2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
