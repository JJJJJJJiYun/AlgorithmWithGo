package main

import "fmt"

type Point struct {
	x, y int
}

func (p Point) test1() {
	fmt.Println(p.x)
}

func (p *Point) test2() {
	fmt.Println(p.y)
}

// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() (x int) {
	x = 1

	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
			x++
		}
	}()
	defer func() {
		panic("test")
	}()
	return x
}
func main() {
	c1 := make(chan int, 1)
	c1 <- 1
	a := <-c1
	fmt.Println(a)
	close(c1)
	b := <-c1
	fmt.Println(b)
}
