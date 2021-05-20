package main

import "fmt"

func main() {
	// 声明一个数组
	var intArray [10]int
	// 此处赋值索引是5，但是是第6个值
	intArray[5] = 111

	fmt.Println(intArray[5])
	// 索引2未赋值，所用默认是对应类型的零值
	fmt.Println(intArray[2])
	// 此处代码会报错，长度是10，取值第11位，数组越界
	//fmt.Println(intArray[10])

}
