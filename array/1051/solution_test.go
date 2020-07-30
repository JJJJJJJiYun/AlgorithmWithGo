package _051

import (
	"fmt"
	"testing"
)

func TestHeightChecker(t *testing.T) {
	heights := []int{1, 1, 4, 2, 1, 3}
	fmt.Println(heightChecker(heights))
}
