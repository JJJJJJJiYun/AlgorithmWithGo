package bilibili

import "math"

// 输入：[[0,3,4,4,5,6,9],[2,4,5,9,9,10],[3,5,8,9,11],[1,4,5,7,9]]
// 输出：5
// 描述：子元素长度不等 内部有序 递增 元素可重复 不唯一 找出在每个数组中都出现的最小值 没有此值返回-1
// 空间复杂度： 0(1) or O(n)

// 实际上需要转化成找公共元素
// 使用hashmap记录元素和在不同数组中出现的次数
// 等于数组个数说明是公共元素
// 找到最小值即可
func Find(mat [][]int) int {
	record := make(map[int]int)
	for _, nums := range mat {
		if len(nums) == 0 {
			return -1
		}
		if len(nums) == 1 {
			record[nums[0]]++
			continue
		}
		i, j := 0, 1
		for i < len(nums) {
			record[nums[i]]++
			for j < len(nums) && nums[j] == nums[i] {
				j++
			}
			i = j
			j++
		}
	}
	min := math.MaxInt32
	for k, v := range record {
		if v == len(mat) && k < min {
			min = k
		}
	}
	return min
}
