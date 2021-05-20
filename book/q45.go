package main

import "fmt"

func main() {
	// defer 可以执行函数，闭包，语句
	defer func() {
		fmt.Println("我是第一个defer")
	}()
	fmt.Println("Println: 1")
	defer fmt.Println("我是第二个defer")
	fmt.Println("Println: 2")
}
