package o12

// 矩阵中的路径

func PathInMatrix(matrix [][]rune, target []rune) bool {
	if matrix == nil || len(matrix) == 0 || matrix[0] == nil || len(matrix[0]) == 0 {
		return false
	}
	res := false
	visited := make([][]bool, 0)
	for _, line := range matrix {
		tmp := make([]bool, 0)
		for range line {
			tmp = append(tmp, false)
		}
		visited = append(visited, tmp)
	}
	var helper func(matrix [][]rune, target []rune, i, j, k int)
	helper = func(matrix [][]rune, target []rune, i, j, k int) {
		if k == len(target) {
			res = true
			return
		}
		if matrix[i][j] != target[k] || visited[i][j] {
			return
		}
		visited[i][j] = true
		// 向上走
		if !res && i-1 >= 0 {
			helper(matrix, target, i-1, j, k+1)
		}
		// 向下走
		if !res && i+1 < len(matrix) {
			helper(matrix, target, i+1, j, k+1)
		}
		// 向左走
		if !res && j-1 >= 0 {
			helper(matrix, target, i, j-1, k+1)
		}
		// 向右走
		if !res && j+1 < len(matrix[0]) {
			helper(matrix, target, i, j+1, k+1)
		}
		visited[i][j] = false
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			helper(matrix, target, i, j, 0)
			if res == true {
				return true
			}
		}
	}
	return false
}
