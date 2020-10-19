package t1267

func countServers(grid [][]int) int {
	if grid == nil || len(grid) == 0 || grid[0] == nil || len(grid[0]) == 0 {
		return 0
	}
	res := 0
	countLine := make([]int, len(grid))
	countCol := make([]int, len(grid[0]))
	for i, line := range grid {
		for j, num := range line {
			if num == 1 {
				countLine[i]++
				countCol[j]++
			}
		}
	}
	for i, line := range grid {
		for j, num := range line {
			if num == 1 && (countLine[i] > 1 || countCol[j] > 1) {
				res++
			}
		}
	}
	return res
}
