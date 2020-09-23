package o2

// 在一个行递增，列也递增的二维数组中找到要找的数

// 每次找到右上角的那个数，这个数是同行最大，同列最小，
// 所以如果这个数比要找的数大，那么可以排除这一列
// 如果这个数比要找的数小，那么可以排除这一行
func FindNumInTwoDimensionArray(nums [][]int, target int) bool {
	if nums == nil || len(nums) == 0 || nums[0] == nil || len(nums[0]) == 0 {
		return false
	}
	i, j := 0, len(nums[0])-1
	for i < len(nums) && j >= 0 {
		if nums[i][j] == target {
			return true
		} else if nums[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

// 最朴实的遍历解法
//func FindNumInTwoDimensionArray(nums [][]int, target int) bool {
//	if nums == nil || len(nums) == 0 || nums[0] == nil || len(nums[0]) == 0 {
//		return false
//	}
//	for i := 0; i < len(nums); i++ {
//		for j := 0; j < len(nums[0]); j++ {
//			if nums[i][j] > target {
//				break
//			}
//			if nums[i][j] == target {
//				return true
//			}
//		}
//	}
//	return false
//}
