package microsoft

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1s := strings.Split(version1, ".")
	v2s := strings.Split(version2, ".")
	if len(v1s) > len(v2s) {
		return -compareVersion(version2, version1)
	}
	i := 0
	for ; i < len(v1s); i++ {
		v1, _ := strconv.ParseInt(v1s[i], 10, 64)
		v2, _ := strconv.ParseInt(v2s[i], 10, 64)
		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}
	for ; i < len(v2s); i++ {
		v2, _ := strconv.ParseInt(v2s[i], 10, 64)
		if v2 > 0 {
			return -1
		}
	}
	return 0
}
