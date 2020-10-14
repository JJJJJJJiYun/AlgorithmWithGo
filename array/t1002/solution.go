package t1002

import "math"

// 维护两个hashmap（因为这里规定了小写字母，所以用数组代替）
// 一个记录在所有字符串中字符出现的最小次数
// 另一个记录当前字符串中字符出现的次数
// 不断的维护这两个hashmap
func commonChars(A []string) []string {
	minRecord := make([]int32, 26)
	for i := range minRecord {
		minRecord[i] = math.MaxInt32
	}
	for _, s := range A {
		record := make([]int32, 26)
		for _, c := range s {
			record[c-'a']++
		}
		for i := range minRecord {
			if minRecord[i] > record[i] {
				minRecord[i] = record[i]
			}
		}
	}
	res := make([]string, 0)
	for k, v := range minRecord {
		for i := int32(0); i < v; i++ {
			res = append(res, string(rune('a'+k)))
		}
	}
	return res
}
