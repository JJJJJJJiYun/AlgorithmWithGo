package t438

// 给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
// 字母异位词指字母相同，但排列不同的字符串。
// 不考虑答案输出的顺序。

// 滑动窗口
func findAnagrams(s string, p string) []int {
	left, right := 0, 0
	sr, pr := []rune(s), []rune(p)
	needMap, getMap := make(map[rune]int), make(map[rune]int)
	valid := 0
	result := make([]int, 0)
	for _, r := range pr {
		needMap[r]++
	}
	for right < len(s) {
		r := sr[right]
		right++
		if _, ok := needMap[r]; ok {
			getMap[r]++
			if getMap[r] == needMap[r] {
				valid++
			}
		}
		for valid == len(needMap) {
			if right-left == len(p) {
				result = append(result, left)
			}
			r := sr[left]
			left++
			if _, ok := needMap[r]; ok {
				if getMap[r] == needMap[r] {
					valid--
				}
				getMap[r]--
			}
		}
	}
	return result
}
