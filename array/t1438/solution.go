package t1438

func longestSubarray(nums []int, limit int) int {
	res := 0
	// 维护一个单调增和一个单调减队列，保存当前窗口内的最小和最大值
	var minQ, maxQ []int
	left := 0
	for right, num := range nums {
		// 将当前值加入到队列中，这里会直接截断后面的部分
		for len(minQ) > 0 && minQ[len(minQ)-1] > num {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, num)
		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < num {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, num)
		// 判断当前窗口的最大最小值是否满足limit，不满足则右移left
		for len(minQ) > 0 && len(maxQ) > 0 && maxQ[0]-minQ[0] > limit {
			// 右移前要先维护两个队列
			if minQ[0] == nums[left] {
				minQ = minQ[1:]
			}
			if maxQ[0] == nums[left] {
				maxQ = maxQ[1:]
			}
			left++
		}
		if res < right-left+1 {
			res = right - left + 1
		}
	}
	return res
}
