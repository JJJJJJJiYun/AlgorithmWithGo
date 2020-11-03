package t941

func validMountainArray(a []int) bool {
	i := 0
	for i < len(a)-1 && a[i] < a[i+1] {
		i++
	}
	if i == 0 || i == len(a)-1 {
		return false
	}
	for i < len(a)-1 && a[i] > a[i+1] {
		i++
	}
	return i == len(a)-1
}
