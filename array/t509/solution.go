package t509

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 || N == 2 {
		return 1
	}
	a, b := 1, 1
	for i := 3; i <= N; i++ {
		a, b = b, a+b
	}
	return b
}
