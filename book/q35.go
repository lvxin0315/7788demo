package main

import "fmt"

func main() {
	testMap1 := make(map[string]int)
	testMap1["one"] = 1
	testMap1["two"] = 2
	fmt.Println("testMap1 len: ", len(testMap1))
}
