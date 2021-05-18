package main

import "fmt"

func main() {
	x := 109
	y := 10
	var z int
	// 普通的赋值符号
	z = x
	fmt.Println("z: ", z)
	// 相加后在赋值
	z += y
	fmt.Println("z += y:  ", z)
	// 相减后在赋值
	z -= y
	fmt.Println("z -= y:  ", z)
	// 相乘后在赋值
	z *= y
	fmt.Println("z *= y:  ", z)
	// 相除后在赋值
	z /= y
	fmt.Println("z /= y:  ", z)
	// 取余后在赋值
	z %= y
	fmt.Println("z %= y:  ", z)
}
