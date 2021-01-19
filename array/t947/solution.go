package t947

func removeStones(stones [][]int) int {
	u := NewUnionSet()
	for _, stone := range stones {
		u.Union(stone[0]+10000, stone[1])
	}
	return len(stones) - u.GetCount()
}

type unionSet struct {
	set   map[int]int
	rank  map[int]int
	count int
}

func NewUnionSet() *unionSet {
	return &unionSet{
		rank:  map[int]int{},
		set:   map[int]int{},
		count: 0,
	}
}

func (u *unionSet) GetCount() int {
	return u.count
}

func (u *unionSet) getRoot(p int) int {
	if _, ok := u.set[p]; !ok {
		u.set[p] = p
		u.count++
	}
	if p != u.set[p] {
		u.set[p] = u.getRoot(u.set[p])
	}
	return u.set[p]
}

func (u *unionSet) Union(p, q int) {
	pRoot := u.getRoot(p)
	qRoot := u.getRoot(q)
	if pRoot != qRoot {
		if u.rank[pRoot] < u.rank[qRoot] {
			u.set[pRoot] = qRoot
		} else if u.rank[qRoot] < u.rank[pRoot] {
			u.set[qRoot] = pRoot
		} else {
			u.set[pRoot] = qRoot
			u.rank[qRoot] += 1
		}
		u.count--
	}
}
