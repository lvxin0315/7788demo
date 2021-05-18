package main

import "fmt"

const A = 1

func main() {
	const B = 2
	// 调用testFunc函数
	testFunc1()
}

func testFunc1() {
	const C = 3
	fmt.Println("A: ", A)
	fmt.Println("B: ", B) // TODO 此处代码会报错
	fmt.Println("C: ", C)
}
