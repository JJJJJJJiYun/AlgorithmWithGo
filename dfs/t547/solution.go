package t547

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	var dfs func(int)
	dfs = func(from int) {
		visited[from] = true
		for to, isCon := range isConnected[from] {
			if isCon == 1 && !visited[to] {
				dfs(to)
			}
		}
	}
	res := 0
	for i, v := range visited {
		if !v {
			res++
			dfs(i)
		}
	}
	return res
}
