package o33

func IsPostOrderTraversal(t []int) bool {
	return helper(t, 0, len(t)-1)
}

func helper(t []int, s, e int) bool {
	if s >= e {
		return true
	}
	root := t[e]
	i := s
	for ; i < e; i++ {
		if t[i] > root {
			break
		}
	}
	index := i
	for ; i < e; i++ {
		if t[i] < root {
			return false
		}
	}
	return helper(t, s, index-1) && helper(t, index, e-1)
}
