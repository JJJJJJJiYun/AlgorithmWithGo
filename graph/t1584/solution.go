package t1584

import "sort"

func minCostConnectPoints(points [][]int) int {
	ans := 0
	u := NewUnionSet(len(points))
	n := len(points)
	type edge struct {
		v, w, dis int
	}
	var edges []edge
	for i, p := range points {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{i, j, dist(p, points[j])})
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].dis < edges[j].dis })
	left := n - 1
	for _, e := range edges {
		if u.Union(e.v, e.w) {
			ans += e.dis
			left--
			if left == 0 {
				break
			}
		}
	}
	return ans
}

func dist(p, q []int) int {
	return abs(p[0]-q[0]) + abs(p[1]-q[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type unionSet struct {
	set  []int
	rank []int
}

func NewUnionSet(n int) *unionSet {
	u := &unionSet{
		rank: make([]int, n),
		set:  make([]int, n),
	}
	for i := 0; i < n; i++ {
		u.set[i] = i
	}
	return u
}

func (u *unionSet) getRoot(p int) int {
	if p != u.set[p] {
		u.set[p] = u.getRoot(u.set[p])
	}
	return u.set[p]
}

func (u *unionSet) Union(p, q int) bool {
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
		return true
	}
	return false
}
