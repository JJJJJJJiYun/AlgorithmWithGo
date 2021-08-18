package sort

const (
	a = 1
)

var (
	b = [a]int{}
)

func QuickSort(nums []int) {
	quickSortHelper(nums, 0, len(nums)-1)
}

func quickSortHelper(nums []int, left, right int) {
	if left < right {
		index := getIndex(nums, left, right)
		quickSortHelper(nums, left, index-1)
		quickSortHelper(nums, index+1, right)
	}
}

func getIndex(nums []int, left, right int) int {
	pivot := nums[left]
	for left < right {
		for left < right && pivot <= nums[right] {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}
