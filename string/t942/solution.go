package t942

// DIIDD
// 501432
func diStringMatch(S string) []int {
	count := 0
	for _, c := range S {
		if c == 'I' {
			count++
		}
	}
	res := make([]int, len(S)+1)
	min, max := count, count
	res[len(res)-1] = count
	for i := len(res) - 2; i >= 0; i-- {
		if S[i] == 'D' {
			max++
			res[i] = max
		} else {
			min--
			res[i] = min
		}
	}
	return res
}

func diStringMatchOfficial(S string) []int {
	res := make([]int, len(S)+1)
	min, max := 0, len(S)
	for i := range res {
		if S[i] == 'D' {
			res[i] = max
			max--
		} else {
			res[i] = min
			min++
		}
	}
	res[len(S)] = min
	return res
}
