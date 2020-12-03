package t1492

func kthFactor(n int, k int) int {
	factor := 1
	count := 0
	for ; factor*factor <= n; factor++ {
		if n%factor == 0 {
			count++
			if count == k {
				return factor
			}
		}
	}
	factor--
	if factor*factor == n {
		factor--
	}
	for ; factor >= 1; factor-- {
		if n%factor == 0 {
			count++
			if count == k {
				return n / factor
			}
		}
	}
	return -1
}
