package main

import "fmt"

func main() {
	funcDemo5(100, 200, 201, 300)
}

// 函数例子4 变量个数不定
func funcDemo5(args ...int) {
	fmt.Println("我是funcDemo5的代码")
	for i, arg := range args {
		fmt.Println(fmt.Sprintf("第%d个参数，值是：%v", i, arg))
	}
}
