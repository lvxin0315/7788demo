package main

import "fmt"

func main() {
	// 声明一个slice
	var intList []int
	fmt.Println(intList)
	// 声明一个slice,同时赋值
	var intList1 = [10]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	fmt.Println(intList1)
	// 声明一个slice,同时指定索引赋值
	var intList2 = []int{1: 11, 2: 12, 9: 19}
	fmt.Println(intList2)
}
