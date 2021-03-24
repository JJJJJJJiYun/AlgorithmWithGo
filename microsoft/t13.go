package microsoft

var romanToIntMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) (res int) {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if i+1 < len(s) && ((c == 'I' && (s[i+1] == 'V' || s[i+1] == 'X')) ||
			(c == 'X' && (s[i+1] == 'L' || s[i+1] == 'C')) ||
			(c == 'C' && (s[i+1] == 'D' || s[i+1] == 'M'))) {
			res += romanToIntMap[s[i+1]] - romanToIntMap[c]
			i++
			continue
		}
		res += romanToIntMap[c]
	}
	return
}
