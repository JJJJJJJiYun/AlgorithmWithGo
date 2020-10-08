package sort

func HeapSort(nums []int) {
	buildMaxHeap(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums[:i], 0)
	}
}

func buildMaxHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(nums, i)
	}
}

func heapify(nums []int, i int) {
	left, right := 2*i+1, 2*i+2
	largest := i
	if left < len(nums) && nums[left] > nums[largest] {
		largest = left
	}
	if right < len(nums) && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		heapify(nums, largest)
	}
}
