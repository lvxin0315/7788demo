package main

import "fmt"

//无参数的函数
func PDemo1() {
	fmt.Println("p demo 1")
}

//带参数函数
func PDemo2(param string) {
	fmt.Println("p demo 2 param:" + param)
}

//变量1
var PDemoVar3 = 0

//输出变量1的函数
func PDemo3() {
	fmt.Println(fmt.Sprintf("p demo %d", PDemoVar3))
}
