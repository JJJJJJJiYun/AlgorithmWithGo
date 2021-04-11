package microsoft

func groupAnagrams(strs []string) [][]string {
	m1 := make(map[[26]int][]string)
	for _, s := range strs {
		temp := [26]int{}
		for _, c := range s {
			temp[c-'a']++
		}
		m1[temp] = append(m1[temp], s)
	}
	res := make([][]string, 0)
	for _, ss := range m1 {
		res = append(res, ss)
	}
	return res
}
