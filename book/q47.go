package main

import "fmt"

func main() {
	// 我们使用defer来让代码执行后，捕获错误
	defer func() {
		errorMessage := recover()
		fmt.Println("errorMessage: ", errorMessage)
	}()
	panic("我是error")
	// 到此就停了,后面的内容执行不到
	fmt.Println("啦啦啦啦啦啦啦啦")
}
