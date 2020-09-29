package o29

import "fmt"

// 顺时针打印矩阵

func PrintMatrix(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 || matrix[0] == nil || len(matrix[0]) == 0 {
		return
	}
	helper(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1)
}

func helper(matrix [][]int, rs, re, cs, ce int) {
	if rs > re || cs > ce {
		return
	}
	m := make([][]int, 0)
	for i := rs; i <= re; i++ {
		temp := make([]int, 0)
		for j := cs; j <= ce; j++ {
			temp = append(temp, matrix[i][j])
		}
		m = append(m, temp)
	}
	rows, cols := len(m), len(m[0])
	for i := 0; i < cols; i++ {
		fmt.Printf("%v ", m[0][i])
	}
	for i := 1; i <= rows-1; i++ {
		fmt.Printf("%v ", m[i][cols-1])
	}
	if rows > 1 {
		for i := 1; i <= cols-1; i++ {
			fmt.Printf("%v ", m[rows-1][cols-1-i])
		}
	}
	if cols > 1 {
		for i := 1; i <= rows-2; i++ {
			fmt.Printf("%v ", m[rows-1-i][0])
		}
	}
	helper(matrix, rs+1, re-1, cs+1, ce-1)
}
