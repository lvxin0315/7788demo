package main

import "fmt"

func main() {
	panic("我是error")
	// 到此就停了,后面的内容执行不到
	fmt.Println("啦啦啦啦啦啦啦啦")
}
