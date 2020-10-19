package t1576

func modifyString(s string) string {
	res := make([]rune, len(s))
	for i, c := range s {
		if c != '?' {
			res[i] = c
		} else {
			var before, after rune
			if i == 0 {
				before = '*'
			} else {
				before = res[i-1]
			}
			if i == len(s)-1 {
				after = '*'
			} else {
				after = rune(s[i+1])
			}
			res[i] = getC(before, after)
		}
	}
	return string(res)
}

func getC(before, after rune) rune {
	for i := int32(0); i < 26; i++ {
		c := 'a' + i
		if c != before && c != after {
			return c
		}
	}
	return 0
}
