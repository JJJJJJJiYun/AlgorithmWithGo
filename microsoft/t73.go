package microsoft

func setZeroes(matrix [][]int) {
	var flag bool
	for _, row := range matrix {
		if row[0] == 0 {
			flag = true
		}
		for i := 1; i < len(row); i++ {
			if row[i] == 0 {
				matrix[0][i] = 0
				row[0] = 0
			}
		}
	}
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
		if flag {
			matrix[i][0] = 0
		}
	}
}
