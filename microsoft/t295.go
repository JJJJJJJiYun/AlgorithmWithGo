package microsoft

import "container/heap"

type MaxIntHeap []int

func (m MaxIntHeap) Len() int            { return len(m) }
func (m MaxIntHeap) Less(i, j int) bool  { return m[i] > m[j] }
func (m MaxIntHeap) Swap(i, j int)       { m[i], m[j] = m[j], m[i] }
func (m MaxIntHeap) Top() int            { return m[0] }
func (m *MaxIntHeap) Push(x interface{}) { *m = append(*m, x.(int)) }
func (m *MaxIntHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

type MinIntHeap []int

func (m MinIntHeap) Len() int            { return len(m) }
func (m MinIntHeap) Less(i, j int) bool  { return m[i] < m[j] }
func (m MinIntHeap) Swap(i, j int)       { m[i], m[j] = m[j], m[i] }
func (m MinIntHeap) Top() int            { return m[0] }
func (m *MinIntHeap) Push(x interface{}) { *m = append(*m, x.(int)) }
func (m *MinIntHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

type MedianFinder struct {
	max *MaxIntHeap
	min *MinIntHeap
}

// MedianFinderConstructor /** initialize your data structure here. */
func MedianFinderConstructor() MedianFinder {
	return MedianFinder{
		max: &MaxIntHeap{},
		min: &MinIntHeap{},
	}
}

func (m *MedianFinder) AddNum(num int) {
	heap.Push(m.max, num)
	heap.Push(m.min, heap.Pop(m.max))
	if m.max.Len() < m.min.Len() {
		heap.Push(m.max, heap.Pop(m.min))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.max.Len() > m.min.Len() {
		return float64(m.max.Top())
	}
	return (float64(m.max.Top()) + float64(m.min.Top())) / 2
}
