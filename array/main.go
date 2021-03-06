package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 9, 9, 8, 0, 7, 1, 7}
	quickSort(nums)
	fmt.Println(nums)
}

func quickSort(nums []int) {
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
	temp := nums[left]
	for left < right {
		for left < right && temp <= nums[right] {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= temp {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = temp
	return left
}
