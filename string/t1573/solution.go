package t1573

func numWays(s string) int {
	countOne := 0
	for _, c := range s {
		if c == '1' {
			countOne++
		}
	}
	// 1的个数不能被3整除
	if countOne%3 != 0 {
		return 0
	}
	// 没有1
	if countOne == 0 {
		return ((len(s) - 1) * (len(s) - 2) / 2) % 1000000007
	}
	// 先找到第一组满足条件的1
	averageOne := countOne / 3
	i := 0
	count := 0
	for count != averageOne {
		if s[i] == '1' {
			count++
		}
		i++
	}
	count1 := 0
	for s[i] == '0' {
		count1++
		i++
	}
	count = 0
	for count != averageOne {
		if s[i] == '1' {
			count++
		}
		i++
	}
	count2 := 0
	for s[i] == '0' {
		count2++
		i++
	}
	return ((count1 + 1) * (count2 + 1)) % 1000000007
}
