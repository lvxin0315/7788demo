package main

import "fmt"

func main() {
	x := 109
	y := 10

	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x % y = ", x%y)
	x++ // 由于输出中不能直接进行++运算，所以单独运行
	fmt.Println("x++ : ", x)
	y-- // 同 x++
	fmt.Println("y-- : ", y)
}
