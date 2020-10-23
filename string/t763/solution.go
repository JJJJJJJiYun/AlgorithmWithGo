package t763

// ababcbacadefegdehijhklij
func partitionLabels(S string) []int {
	res := make([]int, 0)
	record := make([]int, 26)
	for i, c := range S {
		record[c-'a'] = i
	}
	start := 0
	maxEnd := 0
	for i, c := range S {
		if maxEnd < record[c-'a'] {
			maxEnd = record[c-'a']
		}
		if i == maxEnd {
			res = append(res, i-start+1)
			start = i + 1
		}
	}
	return res
}
