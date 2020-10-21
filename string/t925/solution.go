package t925

func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		if name[i] == typed[j] {
			i++
			j++
		} else {
			if j > 0 && typed[j] == typed[j-1] {
				j++
				continue
			} else {
				return false
			}
		}
	}
	// 如果是typed遍历完了，但name还没有，说明不是
	if i < len(name) {
		return false
	}
	// 如果是name遍历晚了，要看typed之后的是不是还是同样的
	j++
	for j < len(typed) && typed[j] == typed[j-1] {
		j++
	}
	return j >= len(typed)
}
