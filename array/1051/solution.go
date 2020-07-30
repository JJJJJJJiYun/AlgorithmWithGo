package _051

// 学校在拍年度纪念照时，一般要求学生按照 非递减 的高度顺序排列。
// 请你返回能让所有学生以 非递减 高度排列的最小必要移动人数。
// 注意，当一组学生被选中时，他们之间可以以任何可能的方式重新排序，而未被选中的学生应该保持不动。

// 实际上最容易想到的方法是直接排序后对比，时间复杂度为nlogn
func heightChecker(heights []int) int {
	arr := make([]int, 101)
	count := 0
	// 记录每一个数，出现的次数
	for _, height := range heights {
		arr[height] ++
	}
	// 对于arr来说其实已经排好序了，从下标1开始，如果当前值大于0，
	// 说明这个数还没有取完，去和heights对应下标的数进行比较，
	// 如果不一样则说明需要移动
	for i, j := 1, 0; i < 101; i++ {
		for arr[i] > 0 {
			if heights[j] != i {
				count++
			}
			j++
			arr[i]--
		}
	}
	return count
}
