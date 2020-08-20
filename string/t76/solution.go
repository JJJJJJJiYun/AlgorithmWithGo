package t76

import "math"

// 给你一个字符串 S、一个字符串 T 。请你设计一种算法，可以在 O(n) 的时间复杂度内，从字符串 S 里面找出：包含 T 所有字符的最小子串。

// 滑动窗口思想
func minWindow(s string, t string) string {
	left, right := 0, 0
	start, length := 0, math.MaxInt64
	sr, tr := []rune(s), []rune(t)
	valid := 0
	needMap, getMap := make(map[rune]int), make(map[rune]int)
	for _, r := range tr {
		needMap[r]++
	}
	for right < len(s) {
		// 要进入滑动框的字符
		r := sr[right]
		// 右滑窗口
		right++
		// 进行数据更新
		if _, ok := needMap[r]; ok {
			getMap[r]++
			if getMap[r] == needMap[r] {
				valid++
			}
		}
		// 判断是否可以左滑窗口
		for valid == len(needMap) {
			if right-left < length {
				length = right - left
				start = left
			}
			// 要离开滑动框的字符
			r := sr[left]
			// 左滑窗口
			left++
			// 进行数据更新
			if _, ok := needMap[r]; ok {
				if getMap[r] == needMap[r] {
					valid--
				}
				getMap[r]--
			}
		}
	}
	if length == math.MaxInt64 {
		return ""
	}
	return string(sr[start : start+length])
}
