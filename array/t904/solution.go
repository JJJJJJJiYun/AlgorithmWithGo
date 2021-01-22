package t904

func totalFruit(tree []int) int {
	res := -1
	n1, n2 := -1, -1
	i := 0
	for ; i < len(tree); i++ {
		if n1 == -1 {
			n1 = i
			continue
		}
		if tree[n1] == tree[i] {
			continue
		}
		if n2 == -1 {
			n2 = i
			continue
		}
		if tree[n2] == tree[i] {
			continue
		}
		if i-n1 > res {
			res = i - n1
		}
		n1 = i - 1
		for ; n1 >= 0 && tree[n1] == tree[n1-1]; n1-- {
		}
		n2 = i
	}
	if i-n1 > res {
		return i - n1
	}
	return res
}
