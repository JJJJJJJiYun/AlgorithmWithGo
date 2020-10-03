package o46

import "strconv"

func WayOfTransNums(num int) int {
	nums := strconv.FormatInt(int64(num), 10)
	nn, n := 1, 1
	for i := 2; i <= len(nums); i++ {
		if (nums[i-2]-'0')*10+nums[i-1]-'0' > 9 && (nums[i-2]-'0')*10+nums[i-1]-'0' < 26 {
			temp := n
			n = n + nn
			nn = temp
		} else {
			nn = n
		}
	}
	return n
}
