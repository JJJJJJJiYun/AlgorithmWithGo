package main

import "fmt"

func main() {
	m := make(map[int][2]int)
	a := m[0]
	a[0] = 1
	fmt.Println(a)
}
