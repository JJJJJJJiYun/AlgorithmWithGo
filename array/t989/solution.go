package t989

import "strconv"

func addToArrayForm(num []int, k int) []int {
	num2 := make([]int, 0)
	for _, c := range strconv.FormatInt(int64(k), 10) {
		num2 = append(num2, int(c-'0'))
	}
	if len(num) > len(num2) {
		num, num2 = num2, num
	}
	l1, l2 := len(num), len(num2)
	res := make([]int, l2)
	flag := 0
	for i := 1; i <= l1; i++ {
		n1, n2 := num[l1-i], num2[l2-i]
		temp := n1 + n2 + flag
		if temp > 9 {
			temp -= 10
			flag = 1
		} else {
			flag = 0
		}
		res[l2-i] = temp
	}
	for i := l1 + 1; i <= l2; i++ {
		temp := num2[l2-i] + flag
		if temp > 9 {
			temp -= 10
			flag = 1
		} else {
			flag = 0
		}
		res[l2-i] = temp
	}
	if flag == 1 {
		res = append([]int{1}, res...)
	}
	return res
}
