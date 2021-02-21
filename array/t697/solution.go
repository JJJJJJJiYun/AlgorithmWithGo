package t697

func findShortestSubArray(nums []int) int {
	countMap := make(map[int]int)
	recordMap := make(map[int][]int)
	max := -1
	for i, num := range nums {
		// 出现次数加一
		countMap[num]++
		// 更新最大出现次数
		if max < countMap[num] {
			max = countMap[num]
		}
		// 更新数字出现的起始位置
		if recordMap[num] == nil {
			// 如果是第一次出现
			recordMap[num] = make([]int, 2)
			recordMap[num][0] = i
		}
		recordMap[num][1] = i
	}
	minLen := len(nums) + 1
	for num, count := range countMap {
		if count == max && minLen > recordMap[num][1]-recordMap[num][0]+1 {
			minLen = recordMap[num][1] - recordMap[num][0] + 1
		}
	}
	return minLen
}
