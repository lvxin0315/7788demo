package main

import "fmt"

// 在这里写注释
func main() {
	// 声明后直接赋值，此处类型"string" 可以省略，因为Go可以在赋值的时候自动判断出类型是什么
	var myString string = "Hello World"
	fmt.Println(myString)
}
