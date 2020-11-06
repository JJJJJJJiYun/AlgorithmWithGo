package t1356

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		c1, c2 := oneCount(arr[i]), oneCount(arr[j])
		if c1 < c2 {
			return true
		} else if c1 == c2 {
			return arr[i] < arr[j]
		} else {
			return false
		}
	})
	return arr
}

func oneCount(num int) int {
	count := 0
	s := fmt.Sprintf("%b", num)
	for _, c := range s {
		if c == '1' {
			count++
		}
	}
	return count
}
