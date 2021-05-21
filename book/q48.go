package main

import "fmt"

func main() {
	var x int
	fmt.Println("x的0值指针", &x)
	fmt.Println("x的值", x)
	y := 1
	fmt.Println("y的指针", &y)
	fmt.Println("y的值", y)
	var z *int
	fmt.Println("z的指针", &z)
	fmt.Println("z的值", z)
}
