package t1011

// 二分查找，自己先找最大最小限制
func shipWithinDays(weights []int, D int) int {
	max, sum := 0, 0
	for _, w := range weights {
		sum += w
		if max < w {
			max = w
		}
	}
	left, right := max, sum
	for left <= right {
		mid := left + (right-left)/2
		if canShip(weights, mid, D) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func canShip(weights []int, cap, D int) bool {
	cur := weights[0]
	d := 1
	for i := 1; i < len(weights); i++ {
		if cur+weights[i] <= cap {
			cur += weights[i]
		} else {
			d++
			cur = weights[i]
		}
	}
	return d <= D
}
