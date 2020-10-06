package o51

var (
	helper []int
)

func ReversePair(nums []int) int {
	helper = make([]int, len(nums))
	return counter(nums, 0, len(nums)-1)
}

func counter(nums []int, start, end int) int {
	if start < end {
		mid := start + (end-start)/2
		return counter(nums, start, mid) + counter(nums, mid+1, end) + merge(nums, start, end)
	}
	return 0
}

func merge(nums []int, start, end int) int {
	mid := start + (end-start)/2
	i, j, k := mid, end, end
	count := 0
	for i >= start && j > mid {
		if nums[i] < nums[j] {
			helper[k] = nums[i]
			i--
		} else {
			helper[k] = nums[j]
			count += j - mid
			j--
		}
		k--
	}
	for i >= start {
		helper[k] = nums[i]
		k--
		i--
	}
	for j > mid {
		helper[k] = nums[j]
		k--
		j--
	}
	for start <= end {
		nums[start] = helper[start]
		start++
	}
	return count
}
