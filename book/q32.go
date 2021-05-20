package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	// 注意长度2
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1)
	fmt.Println(slice2)
	// 修改
	slice1[0] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)
}
