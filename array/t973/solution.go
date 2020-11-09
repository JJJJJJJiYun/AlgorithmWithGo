package t973

func kClosest(points [][]int, K int) [][]int {
	distances := make([]int, len(points))
	for i, point := range points {
		distances[i] = point[0]*point[0] + point[1]*point[1]
	}
	i, j := 0, len(points)-1
	for {
		index := getIndex(distances, points, i, j)
		if index == K-1 {
			return points[:K]
		} else if index > K-1 {
			j = index - 1
		} else {
			i = index + 1
		}
	}
}

func getIndex(distances []int, points [][]int, left, right int) int {
	pivot := distances[left]
	pivotPoint := points[left]
	for left < right {
		if left < right && distances[right] >= pivot {
			right--
		}
		distances[left] = distances[right]
		points[left] = points[right]
		if left < right && distances[left] <= pivot {
			left++
		}
		distances[right] = distances[left]
		points[right] = points[left]
	}
	distances[left] = pivot
	points[left] = pivotPoint
	return left
}
