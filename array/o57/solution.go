package o57

func SumS(nums []int, s int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		temp := nums[i] + nums[j]
		if temp == s {
			return []int{nums[i], nums[j]}
		} else if temp > s {
			j--
		} else {
			i++
		}
	}
	return []int{-1, -1}
}

// 思路类似，就不验证了，可能有错误

func ContinuousSumS(s int) [][]int {
	if s < 3 {
		return [][]int{{s}}
	}
	res := make([][]int, 0)
	i, j := 1, 2
	for j <= s/2+1 {
		temp := make([]int, 0)
		sum := 0
		for k := i; k <= j; k++ {
			sum += k
			temp = append(temp, k)
		}
		if sum == s {
			res = append(res, temp)
		} else if sum > s {
			i++
		} else {
			j++
		}
	}
	return res
}
