package main

import "fmt"

func main() {
	// 声明变量，默认声明map方式是空（nil）
	var testMap map[string]int
	// 使用make函数声明
	testMap1 := make(map[string]int)
	// TODO 此处会报错 panic: assignment to entry in nil map
	testMap["one"] = 1
	testMap["two"] = 2
	testMap1["one"] = 1
	testMap1["two"] = 2
	fmt.Println("testMap: ", testMap)
	fmt.Println("testMap1: ", testMap1)
}
