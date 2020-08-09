package t1414

// 给你数字 k，请你返回和为 k 的斐波那契数字的最少数目，其中，每个斐波那契数字都可以被使用多次。

// 贪心算法，实际上很容易想到这么写，但证明比较困难
// 第一步：证明 我们不会选取连续两个斐波那契数。
// 第二步：证明 我们不会选取同一个斐波那契数超过一次。
// 第三步：证明我们需要的结论，即 对于任意给定的 k，我们需要选择不超过 k 的最大斐波那契数。
func findMinFibonacciNumbers(k int) int {
	a, b := 1, 1
	fibo := []int{1, 1}
	for a+b <= k {
		fibo = append(fibo, a+b)
		a, b = b, a+b
	}
	cnt := 0
	for i := len(fibo) - 1; i >= 0; i-- {
		if k >= fibo[i] {
			cnt++
			k -= fibo[i]
		}
	}
	return cnt
}
