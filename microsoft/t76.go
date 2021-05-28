package microsoft

func minWindow(s string, t string) string {
	need := map[byte]int{}
	for i := range t {
		need[t[i]]++
	}
	needCnt := len(t)
	i := 0
	start, end := -1, len(s)
	for j := range s {
		cj := s[j]
		if need[cj] > 0 {
			needCnt--
		}
		need[cj]--
		if needCnt == 0 {
			for needCnt == 0 {
				ci := s[i]
				need[ci]++
				if need[ci] > 0 {
					needCnt++
				}
				i++
			}
			if j-i+2 < end-start+1 {
				start, end = i-1, j
			}
		}
	}
	if start == -1 {
		return ""
	}
	return s[start : end+1]
}
