package t1544

func makeGood(s string) string {
	if len(s) < 2 {
		return s
	}
	r := []rune(s)
	for i := 1; i < len(r); i++ {
		if r[i]-'A'+'a' == r[i-1] || r[i]-'a'+'A' == r[i-1] {
			if i == 1 {
				r = r[2:]
				i = 0
			} else {
				r = append(r[:i-1], r[i+1:]...)
				i = i - 2
			}
		}
	}
	return string(r)
}
