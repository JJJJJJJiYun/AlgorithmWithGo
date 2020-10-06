package o50

func FirstAppearOnceChar(s string) string {
	record := make(map[int32]int)
	for _, c := range s {
		record[c]++
	}
	for _, c := range s {
		if record[c] == 1 {
			return string(c)
		}
	}
	return ""
}
