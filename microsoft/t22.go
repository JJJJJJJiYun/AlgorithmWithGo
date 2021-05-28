package microsoft

func generateParenthesis(n int) (res []string) {
	var s string
	var helper func(open, close int)
	helper = func(open, close int) {
		if len(s) == n*2 {
			res = append(res, s)
			return
		}
		if open < n {
			s += "("
			helper(open+1, close)
			s = s[:len(s)-1]
		}
		if close < open {
			s += ")"
			helper(open, close+1)
			s = s[:len(s)-1]
		}
	}
	helper(0, 0)
	return
}
