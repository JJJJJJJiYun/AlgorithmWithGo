package t1030

// 桶排序
func allCellsDistOrder(n, m, r0, c0 int) [][]int {
	maxDist := max(r0, n-1-r0) + max(c0, m-1-c0)
	buckets := make([][][]int, maxDist+1)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dist := abs(i-r0) + abs(j-c0)
			buckets[dist] = append(buckets[dist], []int{i, j})
		}
	}
	ans := make([][]int, 0, n*m)
	for _, bucket := range buckets {
		ans = append(ans, bucket...)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var move = [][]int{{-1, 1}, {1, 1}, {1, -1}, {-1, -1}}

func allCellsDistOrder2(n, m, r0, c0 int) [][]int {
	ans := make([][]int, 0, n*m)
	ans = append(ans, []int{r0, c0})
	dist := 1
	maxDist := max(r0, n-1-r0) + max(c0, m-1-c0)
	row, col := r0, c0
	for ; dist <= maxDist; dist++ {
		col--
		for _, mo := range move {
			for j := 0; j < dist; j++ {
				if row >= 0 && row < n && col >= 0 && col < m {
					ans = append(ans, []int{row, col})
				}
				row += mo[0]
				col += mo[1]
			}
		}
	}
	return ans
}
