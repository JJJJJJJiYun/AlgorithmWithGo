package t842

import "math"

func splitIntoFibonacci(S string) []int {
	var res []int
	var backtrack func(start int) bool
	backtrack = func(start int) bool {
		if start == len(S) {
			return len(res) > 2
		}
		cur := 0
		for i := start; i < len(S); i++ {
			if S[start] == '0' && i > start {
				break
			}
			cur = cur*10 + int(S[i]-'0')
			if cur >= math.MaxInt32 {
				break
			}
			if len(res) >= 2 {
				s := res[len(res)-1] + res[len(res)-2]
				if cur > s {
					break
				}
				if cur < s {
					continue
				}
			}
			res = append(res, cur)
			if backtrack(i + 1) {
				return true
			}
			res = res[:len(res)-1]
		}
		return false
	}
	backtrack(0)
	return res
}
