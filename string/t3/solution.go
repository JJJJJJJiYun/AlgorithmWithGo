package t3

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	sr := []rune(s)
	recorder := make(map[rune]int)
	res := 0
	for right < len(s) {
		r := sr[right]
		right++
		recorder[r]++
		for recorder[r] > 1 {
			r := sr[left]
			left++
			recorder[r]--
		}
		// 注意更新答案的时间点
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
