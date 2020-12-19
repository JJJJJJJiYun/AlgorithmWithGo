package t48

func rotate(matrix [][]int) {
	start, end := 0, len(matrix)-1
	for start <= end {
		for i := 0; i < end-start; i++ {
			n1, n2, n3, n4 := matrix[start][start+i], matrix[start+i][end], matrix[end][end-i], matrix[end-i][start]
			matrix[start][start+i], matrix[start+i][end], matrix[end][end-i], matrix[end-i][start] = n4, n1, n2, n3
		}
		start++
		end--
	}
}
