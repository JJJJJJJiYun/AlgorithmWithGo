package o38

import (
	"fmt"
	"sort"
)

// 字符串全排序

func StringPermutation(s string) {
	r := []rune(s)
	visited := make([]bool, len(s))
	helper(r, []rune{}, visited)
}

func helper(r []rune, s []rune, visited []bool) {
	if len(s) == len(r) {
		fmt.Println(string(s))
	}
	for i := range r {
		if visited[i] {
			continue
		}
		s = append(s, r[i])
		visited[i] = true
		helper(r, s, visited)
		s = s[:len(s)-1]
		visited[i] = false
	}
}

// 有重复字符的字符串全排序

func StringPermutationWithRepeated(s string) {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	visited := make([]bool, len(s))
	helper2(r, []rune{}, visited)
}

func helper2(r []rune, s []rune, visited []bool) {
	if len(s) == len(r) {
		fmt.Println(string(s))
	}
	for i := range r {
		if visited[i] || (i > 0 && r[i] == r[i-1] && visited[i-1]) {
			continue
		}
		s = append(s, r[i])
		visited[i] = true
		helper2(r, s, visited)
		visited[i] = false
		s = s[:len(s)-1]
	}
}
