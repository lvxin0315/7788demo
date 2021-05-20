package main

import "fmt"

func main() {
	// 声明变量，默认声明map方式是空（nil）
	var testMap map[string]int
	// 使用make函数声明
	testMap1 := make(map[string]int)
	fmt.Println("testMap: ", testMap == nil)
	fmt.Println("testMap1: ", testMap1 == nil)
}
