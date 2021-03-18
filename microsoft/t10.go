package microsoft

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	var firstMatch bool
	if len(s) > 0 {
		firstMatch = p[0] == '.' || p[0] == s[0]
	}
	if len(p) > 1 && p[1] == '*' {
		return isMatch(s, p[2:]) || (len(s) > 0 && firstMatch && isMatch(s[1:], p))
	}
	return len(s) > 0 && firstMatch && isMatch(s[1:], p[1:])
}
