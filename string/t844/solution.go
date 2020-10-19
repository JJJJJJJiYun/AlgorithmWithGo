package t844

func backspaceCompare(S string, T string) bool {
	i, j := len(S)-1, len(T)-1
	skipS, skipT := 0, 0
	// 注意这个结束条件是或
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if S[i] == '#' {
				skipS++
			} else if skipS > 0 {
				skipS--
			} else {
				break
			}
			i--
		}
		for j >= 0 {
			if T[j] == '#' {
				skipT++
			} else if skipT > 0 {
				skipT--
			} else {
				break
			}
			j--
		}
		// 注意这里的判断
		if i >= 0 && j >= 0 {
			if S[i] != T[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}
