package t387

func firstUniqChar(s string) int {
	recorder := make([]int, 26)
	for i := range recorder {
		recorder[i] = -1
	}
	for i, c := range s {
		if recorder[c-'a'] == -1 {
			recorder[c-'a'] = i
		} else {
			recorder[c-'a'] = -2
		}
	}
	res := len(s)
	for i := range recorder {
		if recorder[i] != -2 && recorder[i] != -1 && recorder[i] < res {
			res = recorder[i]
		}
	}
	if res == len(s) {
		return -1
	}
	return res
}
