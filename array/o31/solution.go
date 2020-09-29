package o31

// 入栈和出栈序列是否匹配

func IsPopOfPush(push, pop []int) bool {
	helper := make([]int, 0)
	i := 0
	for _, v := range pop {
		if len(helper) > 0 && helper[len(helper)-1] == v {
			helper = helper[:len(helper)-1]
			continue
		}
		if i >= len(push) {
			return false
		}
		for ; i < len(push); i++ {
			helper = append(helper, push[i])
			if push[i] == v {
				break
			}
		}
		helper = helper[:len(helper)-1]
		i++
	}
	return true
}
