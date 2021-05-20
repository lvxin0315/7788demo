package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	// append操作
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1)
	fmt.Println(slice2)
	// 修改
	slice1[0] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)
}
