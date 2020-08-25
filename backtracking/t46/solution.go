package t46

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	// 对于匿名函数的递归调用，需要先声明变量完成绑定
	var helper func(nums, current []int)
	helper = func(nums, current []int) {
		if len(nums) == 0 {
			// slice 共享底层数组，所以这里要做一个拷贝操作，不能直接加入
			tmp := make([]int, len(current))
			copy(tmp, current)
			result = append(result, tmp)
			return
		}
		for i := range nums {
			current = append(current, nums[i])
			// 这里一定要做一个重新分配数组的动作
			// 如果直接采用 append(nums[:i], nums[i+1:]...) 的方式传入新的nums
			// 会导致底层nums数组被更改
			tmp := make([]int, 0)
			tmp = append(tmp, nums[:i]...)
			tmp = append(tmp, nums[i+1:]...)
			helper(tmp, current)
			current = current[:len(current)-1]
		}
	}
	helper(nums, []int{})
	return result
}
