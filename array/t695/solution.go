package t695

// 给定一个包含了一些 0 和 1 的非空二维数组grid 。
// 一个岛屿是由一些相邻的1(代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。
// 你可以假设grid 的四个边缘都被 0（代表水）包围着。
// 找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)

// 递归，深度优先遍历，注意把grip[i][j]直接赋值成0这一步，因为每一个土地只能够被路径使用一次
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area := getArea(i, j, grid)
				maxArea = max(area, maxArea)
			}
		}
	}
	return maxArea
}

func getArea(i, j int, grid [][]int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return 1 + getArea(i-1, j, grid) + getArea(i+1, j, grid) + getArea(i, j-1, grid) + getArea(i, j+1, grid)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
