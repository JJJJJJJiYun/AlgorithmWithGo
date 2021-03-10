package microsoft

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	var n int64
	var flag bool
	i := 0
	if s[0] == '+' {
		flag = false
		i++
	} else if s[0] >= '0' && s[0] <= '9' {
		flag = false
	} else if s[0] == '-' {
		flag = true
		i++
	} else {
		return 0
	}
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		num := int64(s[i] - '0')
		n = n*10 + num
		if flag && n > math.MaxInt32+1 {
			return math.MinInt32
		} else if !flag && n > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return int(n)
}
