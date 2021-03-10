package microsoft

func numIslands(grid [][]byte) (res int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				res++
				numIslandsHelper(grid, i, j)
			}
		}
	}
	return
}

func numIslandsHelper(grid [][]byte, i, j int) {
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	if i > 0 {
		numIslandsHelper(grid, i-1, j)
	}
	if i < len(grid)-1 {
		numIslandsHelper(grid, i+1, j)
	}
	if j > 0 {
		numIslandsHelper(grid, i, j-1)
	}
	if j < len(grid[0])-1 {
		numIslandsHelper(grid, i, j+1)
	}
}
