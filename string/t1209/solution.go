package t1209

func removeDuplicates(s string, k int) string {
	if len(s) < k {
		return s
	}
	record := make([]int, len(s))
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if i == 0 || r[i] != r[i-1] {
			record[i] = 1
		} else {
			record[i] = record[i-1] + 1
			if record[i] == k {
				r = append(r[:i-k+1], r[i+1:]...)
				i = i - k
			}
		}
	}
	return string(r)
}
