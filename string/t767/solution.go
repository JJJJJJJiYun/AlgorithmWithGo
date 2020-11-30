package t767

func reorganizeString(s string) string {
	record := make([]int, 26)
	max := 0
	for _, c := range s {
		record[c-'a']++
		if record[c-'a'] > max {
			max = record[c-'a']
		}
	}
	if max > (len(s)+1)/2 {
		return ""
	}
	even, odd := 0, 1
	res := make([]byte, len(s))
	for i, count := range record {
		c := byte('a' + i)
		for count > 0 && count <= len(s)/2 && odd < len(s) {
			res[odd] = c
			count--
			odd += 2
		}
		for count > 0 {
			res[even] = c
			count--
			even += 2
		}
	}
	return string(res)
}
