package main

import "fmt"

func main() {
	// 调用
	funcDemo1()
	result2 := funcDemo2(2)
	fmt.Println("我是funcDemo2的返回值：", result2)
	result3 := funcDemo3(2)
	fmt.Println("我是funcDemo3的返回值：", result3)

}

// 函数例子1 无参数，无返回值
func funcDemo1() {
	fmt.Println("我是funcDemo1的代码")
}

// 函数例子2 有参数和返回值，注意入参和出参都需要指定类型
func funcDemo2(param int) int {
	fmt.Println("我是funcDemo2的代码")
	param += 100
	return param
}

// 函数例子3 有参数和返回值，返回值在函数声明时定义
func funcDemo3(param int) (result int) {
	fmt.Println("我是funcDemo3的代码")
	result = param * 2
	// 当在函数声明时定义返回值，所以return可以不用返回变量
	return
}
