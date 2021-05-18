package main

import "fmt"

// 在这里写注释
func main() {
	// 同时声明多个string
	var myStr1, myStr2, myStr3 string

	// 同时给多个变量赋值
	v1, v2, v3 := 1, true, ""

	// 另一种声明多个变量方式
	var (
		i1 = 5
		i2 = 10
		i3 = 15
	)

	fmt.Println("myStr1: ", myStr1, " myStr2: ", myStr2, " myStr3: ", myStr3)

	fmt.Println("v1: ", v1, " v2: ", v2, " v3: ", v3)

	fmt.Println("i1: ", i1, " i2: ", i2, " i3: ", i3)
}
