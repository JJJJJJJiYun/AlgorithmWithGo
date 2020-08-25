package t46

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	helper(nums, []int{}, &result)
	return result
}

func helper(nums, current []int, result *[][]int) {
	if len(nums) == 0 {
		// slice 共享底层数组，所以这里要做一个拷贝操作，不能直接加入
		tmp := make([]int, len(current))
		copy(tmp, current)
		*result = append(*result, tmp)
		return
	}
	for i := range nums {
		current = append(current, nums[i])
		// 这里一定要做一个重新分配数组的动作
		// 如果直接采用 append(nums[:i], nums[i+1:]...)的方式传入新的nums
		// 会导致底层nums数组被更改
		tmp := make([]int, 0)
		tmp = append(tmp, nums[:i]...)
		tmp = append(tmp, nums[i+1:]...)
		helper(tmp, current, result)
		current = current[:len(current)-1]
	}
}
