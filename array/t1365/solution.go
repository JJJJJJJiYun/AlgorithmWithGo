package t1365

// solution1：排序后计数

// solution2：计数排序
func smallerNumbersThanCurrent(nums []int) []int {
	count := make([]int, 101)
	for _, num := range nums {
		count[num]++
	}
	// 小于这个数的个数等于前面的累加
	for i := 0; i < 100; i++ {
		count[i+1] += count[i]
	}
	res := make([]int, 0)
	for _, num := range nums {
		if num > 0 {
			res = append(res, count[num-1])
		} else {
			res = append(res, 0)
		}
	}
	return res
}
