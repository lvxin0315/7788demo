package main

import "fmt"

func main() {
	var intArray = [10]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	// 一个长度为10的数组，索引是 0到9
	for i := 0; i < len(intArray); i++ {
		fmt.Println("索引i: ", i, " 值: ", intArray[i])
	}
}
