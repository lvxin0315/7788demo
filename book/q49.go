package main

import "fmt"

func main() {
	x := 100
	fmt.Println("x: ", x)
	funcDemo8(&x)
	fmt.Println("x: ", x)
}

// 修改指针的值
func funcDemo8(p *int) {
	*p = 200
}
