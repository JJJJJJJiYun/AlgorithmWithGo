package sort

func MergeSort(nums []int) {
	mergeSortHelper(nums, 0, len(nums)-1)
}

func MergeSort2(nums []int) {
	for i := 2; i < len(nums); i += i {
		left := 0
		right := left + i - 1
		for right < len(nums) {
			merge(nums, left, left+(right-left)/2, right)
			left += i
			right += i
		}
		if left < len(nums) {
			merge(nums, left-i, left-1, len(nums)-1)
		}
	}
}

func mergeSortHelper(nums []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		mergeSortHelper(nums, left, mid)
		mergeSortHelper(nums, mid+1, right)
		merge(nums, left, left+(right-left)/2, right)
	}
}

func merge(nums []int, left, mid, right int) {
	helper := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			helper[k] = nums[i]
			i++
		} else {
			helper[k] = nums[j]
			j++
		}
		k++
	}
	for i <= mid {
		helper[k] = nums[i]
		i++
		k++
	}
	for j <= right {
		helper[k] = nums[j]
		j++
		k++
	}
	k = 0
	for left <= right {
		nums[left] = helper[k]
		k++
		left++
	}
}
