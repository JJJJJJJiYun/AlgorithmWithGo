package o5

import "strings"

// 替换空格
func ReplaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}
