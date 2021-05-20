package main

import "fmt"

func main() {
	result1, result2, result3 := funcDemo4()
	fmt.Println("我是funcDemo4的返回值：", result1, " ", result2, " ", result3)

}

// 函数例子4 多返回值
func funcDemo4() (int, int, int) {
	fmt.Println("我是funcDemo4的代码")
	return 1, 2, 3
}
