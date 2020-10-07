package o53

func CountNumInSortedArray(nums []int, num int) int {
	return findRightBorder(nums, num) - findLeftBorder(nums, num) + 1
}

func findLeftBorder(nums []int, num int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] < num {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	if i >= 0 && nums[i] == num {
		return i
	}
	return -1
}

func findRightBorder(nums []int, num int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] <= num {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	if i-1 <= len(nums)-1 && nums[i-1] == num {
		return i - 1
	}
	return -1
}

// 0,1,2,3,5
// 0,2,3,4,5
// 1
func MissingNum(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == mid {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}

// -3,1,5,6,7
func FindNumMatchIndex(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == mid {
			return mid
		} else if nums[mid] > mid {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return -1
}
