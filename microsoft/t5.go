package microsoft

func longestPalindrome(s string) string {
	start, end := 0, 0
	for i := range s {
		left1, right1 := expend(s, i, i)
		left2, right2 := expend(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start+1 : end]
}

func expend(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left, right
}
