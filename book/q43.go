package main

import "fmt"

func main() {
	// 存放闭包
	demo6 := funcDemo6()
	fmt.Println("funcDemo6: ", demo6())
}

// 函数6
func funcDemo6() func() string {
	fmt.Println("我是funcDemo6的代码")
	return func() string {
		return "demo6"
	}
}
