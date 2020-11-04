package t1557

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	record := make(map[int]struct{})
	res := make([]int, 0)
	for _, edge := range edges {
		record[edge[1]] = struct{}{}
	}
	for i := 0; i < n; i++ {
		if _, ok := record[i]; !ok {
			res = append(res, i)
		}
	}
	return res
}
