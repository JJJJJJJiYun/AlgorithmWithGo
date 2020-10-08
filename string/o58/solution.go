package o58

func reverseString(s string) string {
	r := []rune(s)
	i, j := 0, len(r)-1
	for i < j {
		temp := r[i]
		r[i] = r[j]
		r[j] = temp
		i++
		j--
	}
	return string(r)
}
