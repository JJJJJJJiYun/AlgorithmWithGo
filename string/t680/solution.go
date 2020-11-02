package t680

func validPalindrome(s string) bool {
	if len(s) <= 2 {
		return true
	}
	i, j := 0, len(s)-1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return helper(s, i+1, j) || helper(s, i, j-1)
		}
	}
	return true
}

func helper(s string, start, end int) bool {
	for start < end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			return false
		}
	}
	return true
}
