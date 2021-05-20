package main

import "fmt"

func main() {
	var intArray = [10]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	intList := intArray[2:6]
	fmt.Println("intArray len: ", len(intArray))
	fmt.Println("intList len: ", len(intList))
}
