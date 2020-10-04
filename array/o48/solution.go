package o48

func LengthOfMaxSubStringWithNoRepeated(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	max := 0
	i, j := 0, 0
	record := make(map[uint8]int)
	for j < len(s) {
		if v, ok := record[s[j]]; ok && i <= v {
			i = record[s[j]] + 1
		}
		if j-i+1 > max {
			max = j - i + 1
		}
		record[s[j]] = j
		j++
	}
	return max
}
