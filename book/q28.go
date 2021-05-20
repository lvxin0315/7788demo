package main

import "fmt"

func main() {
	// 声明一个长度10的slice
	var intList = make([]int, 10)
	fmt.Println(intList)
	// 声明一个长度5的slice，占用空间6
	var intList1 = make([]int, 5, 6)
	fmt.Println(intList1)
}
