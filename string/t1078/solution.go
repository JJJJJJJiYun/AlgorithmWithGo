package t1078

import "strings"

func findOcurrences(text string, first string, second string) []string {
	splits := strings.Split(text, " ")
	res := make([]string, 0)
	if len(splits) < 3 {
		return res
	}
	i, j := 0, 1
	for j <= len(splits)-2 {
		if splits[i] == first && splits[j] == second {
			res = append(res, splits[j+1])
		}
		i++
		j++
	}
	return res
}
