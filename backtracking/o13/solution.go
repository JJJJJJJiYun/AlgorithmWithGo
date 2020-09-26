package o13

import "strconv"

// 机器人的运动范围

func RobotMove(rows, cols, threshold int) int {
	var helper func(rows, cols, threshold, i, j int)
	res := 0
	visited := make([][]bool, 0)
	for i := 0; i < rows; i++ {
		tmp := make([]bool, 0)
		for j := 0; j < cols; j++ {
			tmp = append(tmp, false)
		}
		visited = append(visited, tmp)
	}
	helper = func(rows, cols, threshold, i, j int) {
		if i < 0 || i >= rows || j < 0 || j >= cols || getNumSum(i)+getNumSum(j) > threshold || visited[i][j] {
			return
		}
		res++
		visited[i][j] = true
		helper(rows, cols, threshold, i-1, j)
		helper(rows, cols, threshold, i+1, j)
		helper(rows, cols, threshold, i, j-1)
		helper(rows, cols, threshold, i, j+1)
	}
	helper(rows, cols, threshold, 0, 0)
	return res
}

func getNumSum(num int) int {
	numStr := strconv.FormatInt(int64(num), 10)
	sum := int32(0)
	for _, s := range numStr {
		sum += s - '0'
	}
	return int(sum)
}
