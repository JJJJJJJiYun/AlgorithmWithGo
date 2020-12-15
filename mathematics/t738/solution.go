package t738

import "strconv"

func monotoneIncreasingDigits(n int) int {
	bytesN := []byte(strconv.FormatInt(int64(n), 10))
	if len(bytesN) == 1 {
		return n
	}
	i := 1
	for i < len(bytesN) && bytesN[i] >= bytesN[i-1] {
		i++
	}
	if i == len(bytesN) {
		return n
	}
	i--
	for i >= 1 && bytesN[i] == bytesN[i-1] {
		i--
	}
	bytesN[i]--
	i++
	for ; i < len(bytesN); i++ {
		bytesN[i] = '9'
	}
	num, _ := strconv.ParseInt(string(bytesN), 10, 64)
	return int(num)
}
