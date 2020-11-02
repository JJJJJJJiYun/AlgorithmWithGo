package t463

func islandPerimeter(grid [][]int) int {
	count := 0
	for i, line := range grid {
		for j, v := range line {
			if v == 1 {
				// 上
				if i == 0 || grid[i-1][j] == 0 {
					count++
				}
				// 下
				if i == len(grid)-1 || grid[i+1][j] == 0 {
					count++
				}
				// 左
				if j == 0 || grid[i][j-1] == 0 {
					count++
				}
				// 右
				if j == len(grid[0])-1 || grid[i][j+1] == 0 {
					count++
				}
			}
		}
	}
	return count
}
