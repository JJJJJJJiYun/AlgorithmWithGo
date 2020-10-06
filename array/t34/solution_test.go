package t34

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	fmt.Println(searchRange([]int{1, 1, 2, 3, 3, 3, 4, 5, 5}, 3))
}
