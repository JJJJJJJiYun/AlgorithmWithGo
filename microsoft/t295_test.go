package microsoft

import (
	"fmt"
	"testing"
)

func TestMedianFinderConstructor(t *testing.T) {
	c := MedianFinderConstructor()
	c.AddNum(1)
	c.AddNum(2)
	fmt.Println(c.FindMedian())
	c.AddNum(3)
	fmt.Println(c.FindMedian())

}
