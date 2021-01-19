package t1018

func prefixesDivBy5(a []int) []bool {
	ans := make([]bool, len(a))
	x := 0
	for i, v := range a {
		x = (x<<1 | v) % 5
		ans[i] = x == 0
	}
	return ans
}
