package t832

func flipAndInvertImage(matrix [][]int) [][]int {
	n := len(matrix)
	for i := range matrix {
		j, k := 0, n-1
		for j < k {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
			matrix[i][j] = reverseNum(matrix[i][j])
			matrix[i][k] = reverseNum(matrix[i][k])
			j++
			k--
		}
		if j == k {
			matrix[i][j] = reverseNum(matrix[i][j])
		}
	}
	return matrix
}

func reverseNum(num int) int {
	if num == 1 {
		return 0
	}
	return 1
}
