package microsoft

import "strconv"

func compress(chars []byte) int {
	i := len(chars) - 1
	j := len(chars) - 1
	for i >= 0 {
		c := chars[i]
		count := 1
		for ; i >= 1 && chars[i] == chars[i-1]; i-- {
			count++
		}
		if count == 1 {
			chars[j] = c
			j--
			i--
			continue
		}
		countStr := strconv.FormatInt(int64(count), 10)
		countStrLen := len(countStr)
		for k := countStrLen - 1; k >= 0; k-- {
			chars[j] = countStr[k]
			j--
		}
		chars[j] = c
		j--
		i--
	}
	i = 0
	j++
	for j < len(chars) {
		chars[i] = chars[j]
		i++
		j++
	}
	return i
}
