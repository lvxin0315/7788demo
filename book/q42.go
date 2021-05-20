package main

import "fmt"

func main() {
	// add变量中放的是函数，高端吧
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))
}
