package microsoft

func getPairs(n int) (res []string) {
	var s string
	var helper func(open, close int)
	helper = func(open, close int) {
		// 当达到数量要求时返回
		if len(s) == n*2 {
			res = append(res, s)
			return
		}
		// 左括号小于数量
		if open < n {
			s += "("
			helper(open+1, close)
			s = s[:len(s)-1]
		}
		// 右括号小于左括号
		if close < open {
			s += ")"
			helper(open, close+1)
			s = s[:len(s)-1]
		}
	}
	helper(0, 0)
	return
}
