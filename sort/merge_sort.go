package sort

func MergeSort(nums []int) {
	helper := make([]int, len(nums))
	mergeSortHelper(nums, helper, 0, len(nums)-1)
}

func mergeSortHelper(nums, helper []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		mergeSortHelper(nums, helper, left, mid)
		mergeSortHelper(nums, helper, mid+1, right)
		merge(nums, helper, left, mid, right)
	}
}

func merge(nums, helper []int, left, mid, right int) {
	i, j, k := left, mid+1, left
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
	for left <= right {
		nums[left] = helper[left]
		left++
	}
}

func MergeSort2(nums []int) {
	helper := make([]int, len(nums))
	for i := 1; i < len(nums); i += i {
		left := 0
		for left+2*i-1 < len(nums) {
			merge(nums, helper, left, left+i-1, left+2*i-1)
			left += 2 * i
		}
		// 如果最后还剩下超过步长的元素，那么这些元素merge一次
		if left < len(nums)-i {
			merge(nums, helper, left, left+i-1, len(nums)-1)
		}
	}
}
