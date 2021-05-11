package microsoft

import "sort"

func merge(intervals [][]int) (res [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= right {
			if intervals[i][1] > right {
				right = intervals[i][1]
			}
			continue
		}
		res = append(res, []int{left, right})
		left, right = intervals[i][0], intervals[i][1]
	}
	if len(res) == 0 || left != res[len(res)-1][0] {
		res = append(res, []int{left, right})
	}
	return
}
