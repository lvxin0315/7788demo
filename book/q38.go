package main

import "fmt"

func main() {
	// array
	var intArray = [5]int{10, 11, 12, 13, 14}
	// 遍历intArray i是索引，v是索引对应的值
	for i, v := range intArray {
		fmt.Println("intArray索引i: ", i, " 值: ", v)
	}
	// slice
	var intList = []int{20, 21, 22, 23, 24}
	// 遍历intList i是索引，v是索引对应的值
	for i, v := range intList {
		fmt.Println("intList索引i: ", i, " 值: ", v)
	}
	// map
	testMap1 := make(map[string]int)
	testMap1["one"] = 1
	testMap1["two"] = 2
	// 遍历testMap1 k是键，v是键对应的值
	for k, v := range intList {
		fmt.Println("testMap1键k: ", k, " 值: ", v)
	}
}
