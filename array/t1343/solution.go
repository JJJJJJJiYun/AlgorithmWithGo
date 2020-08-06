package t1343

// 给你一个整数数组 arr 和两个整数 k 和 threshold 。
// 请你返回长度为 k 且平均值大于等于 threshold 的子数组数目。
// 滑动窗口思想
func numOfSubarrays(arr []int, k int, threshold int) int {
	// 注意这里如果是用除法计算，就需要化成float类型，这样用乘积作为目标会简单很多
	count, sum, target := 0, 0, threshold*k
	// 注意这个for循环的写法
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if i < k-1 {
			continue
		}
		if i >= k {
			sum -= arr[i-k]
		}
		if sum >= target {
			count++
		}
	}
	return count
}
