package _343

import (
	"fmt"
	"testing"
)

func TestNumOfSubarrays(t *testing.T) {
	arr := []int{4, 4, 4, 4}
	k := 4
	threshold := 1
	fmt.Println(numOfSubarrays(arr, k, threshold))
}
