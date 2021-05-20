package main

import "fmt"

func main() {
	var intArray = [10]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	// 此处，取出数组中的索引2到6对应值，12，13，14，15
	intList := intArray[2:6]
	fmt.Println(intList)
	// 我们修改数组中的部分值
	intArray[2] = 1
	intArray[3] = 1
	// 同时slice中也变了
	fmt.Println(intList)
	// 我们试着修改一下slice中的值，索引3 对应15
	intList[3] = 50
	fmt.Println(intList)
	fmt.Println(intArray)
}
