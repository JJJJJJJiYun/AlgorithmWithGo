package t567

func checkInclusion(s1 string, s2 string) bool {
	left, right := 0, 0
	r1, r2 := []rune(s1), []rune(s2)
	needMap, getMap := make(map[rune]int), make(map[rune]int)
	valid := 0
	for _, r := range r1 {
		needMap[r]++
	}
	for right < len(s2) {
		r := r2[right]
		right++
		if _, ok := needMap[r]; ok {
			getMap[r]++
			if getMap[r] == needMap[r] {
				valid++
			}
		}
		for valid == len(needMap) {
			if right-left == len(s1) {
				return true
			}
			r := r2[left]
			left++
			if _, ok := needMap[r]; ok {
				if getMap[r] == needMap[r] {
					valid--
				}
				getMap[r]--
			}
		}
	}
	return false
}
