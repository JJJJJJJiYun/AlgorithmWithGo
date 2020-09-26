package o10

// 斐波那契数列

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	a, b := 1, 1
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// 青蛙跳台阶

// dp[n] = dp[n-1] + dp[n-2]
// 实际上就是斐波那契额
func frog(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	a, b := 1, 1
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
