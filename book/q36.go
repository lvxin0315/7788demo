package main

import "fmt"

func main() {
	testMap1 := make(map[string]int)
	testMap1["one"] = 1
	testMap1["two"] = 2
	// 删除
	delete(testMap1, "one")
	fmt.Println("testMap1: ", testMap1)
	fmt.Println("testMap1 len: ", len(testMap1))
}
