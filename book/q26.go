package main

import "fmt"

func main() {
	// 声明一个数组,同时赋值
	var intArray = [10]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	fmt.Println(intArray)
	// 声明一个数组,同时指定索引赋值
	var intArray1 = [10]int{1: 11, 2: 12, 9: 19}
	fmt.Println(intArray1)
}
