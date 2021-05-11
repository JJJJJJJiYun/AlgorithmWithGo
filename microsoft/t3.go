package microsoft

func lengthOfLongestSubstring(s string) int {
	max := 0
	recorder := make(map[rune]int)
	start := 0
	for i, c := range s {
		if index, ok := recorder[c]; ok && start <= index {
			start = index + 1
		} else {
			if i-start+1 > max {
				max = i - start + 1
			}
		}
		recorder[c] = i
	}
	return max
}
