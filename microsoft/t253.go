package microsoft

import "sort"

func minMeetingRooms(intervals [][]int) (res int) {
	var starts, ends []int
	for _, interval := range intervals {
		starts = append(starts, interval[0])
		ends = append(ends, interval[1])
	}
	sort.Ints(starts)
	sort.Ints(ends)
	var i, j int
	for i < len(starts) && j < len(ends) {
		if starts[i] < ends[j] {
			res++
		} else {
			j++
		}
		i++
	}
	return res
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minMeetingRooms2(intervals [][]int) (res int) {
	h := &IntHeap{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for _, interval := range intervals {
		if h.Len() == 0 {
			h.Push(interval[1])
			continue
		}
		minEnd := h.Pop().(int)
		if minEnd > interval[0] {
			h.Push(minEnd)
		}
		h.Push(interval[1])
	}
	return h.Len()
}
