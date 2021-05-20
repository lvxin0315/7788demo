package main

import "fmt"

func main() {
	fmt.Println("funcDemo7: ", funcDemo7(5))
}

// 函数7
func funcDemo7(x int) int {
	fmt.Println("当前x是: ", x)
	if x == 0 {
		return 1
	}
	return x * funcDemo7(x-1)
}
