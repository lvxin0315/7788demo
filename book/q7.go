package main

import "fmt"

// 在所有函数外部声明的变量
var a = 1

func main() {
	// 调用testFunc函数
	testFunc()
}

func testFunc() {
	// 在testFunc 函数中声明的变量
	var b = 2
	fmt.Println("a: ", a) // a是第六行的变量
	fmt.Println("b: ", b)
}
