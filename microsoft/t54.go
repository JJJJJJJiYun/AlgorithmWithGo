package microsoft

func spiralOrder(matrix [][]int) (res []int) {
	return spiralOrderHelper(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1, res)
}

func spiralOrderHelper(matrix [][]int, startLine, endLine, startCol, endCol int, res []int) []int {
	if startLine > endLine || startCol > endCol {
		return res
	}
	if startLine == endLine {
		for i := startCol; i <= endCol; i++ {
			res = append(res, matrix[startLine][i])
		}
		return res
	}
	if startCol == endCol {
		for i := startLine; i <= endLine; i++ {
			res = append(res, matrix[i][startCol])
		}
		return res
	}
	for i := startCol; i <= endCol; i++ {
		res = append(res, matrix[startLine][i])
	}
	for i := startLine + 1; i <= endLine; i++ {
		res = append(res, matrix[i][endCol])
	}
	for i := endCol - 1; i >= startCol; i-- {
		res = append(res, matrix[endLine][i])
	}
	for i := endLine - 1; i >= startLine+1; i-- {
		res = append(res, matrix[i][startCol])
	}
	res = spiralOrderHelper(matrix, startLine+1, endLine-1, startCol+1, endCol-1, res)
	return res
}
