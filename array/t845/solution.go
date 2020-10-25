package t845

// 遍历找就好
func longestMountain(array []int) int {
	if len(array) < 3 {
		return 0
	}
	cur := 1
	max := 0
	for cur < len(array) {
		if array[cur] > array[cur-1] {
			i := cur
			for i < len(array) && array[i] > array[i-1] {
				i++
			}
			if i == len(array) {
				return max
			}
			for i < len(array) && array[i] < array[i-1] {
				i++
			}
			if array[i-1] < array[i-2] && max < i-cur+1 {
				max = i - cur + 1
				cur = i - 1
			}
		}
		cur++

	}
	return max
}

// 官方题解，思路类似，写法不同
func longestMountain2(a []int) (ans int) {
	n := len(a)
	left := 0
	for left+2 < n {
		right := left + 1
		if a[left] < a[left+1] {
			for right+1 < n && a[right] < a[right+1] {
				right++
			}
			if right < n-1 && a[right] > a[right+1] {
				for right+1 < n && a[right] > a[right+1] {
					right++
				}
				if right-left+1 > ans {
					ans = right - left + 1
				}
			} else {
				right++
			}
		}
		left = right
	}
	return
}
