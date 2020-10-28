package t1207

func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int)
	for _, v := range arr {
		count[v]++
	}
	times := make(map[int]struct{})
	for _, c := range count {
		times[c] = struct{}{}
	}
	return len(times) == len(count)
}
