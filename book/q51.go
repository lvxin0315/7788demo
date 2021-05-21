package main

import "fmt"

// 定义一个结构
type demoStruct struct {
	// public
	Name   string // 姓名
	Mobile string // 手机号
	Gender int    // 性别，0-未知，1-男，2-女
	// private
	code string // 内部编码
}

func main() {
	// 初始方式：new 返回的是指针
	x := new(demoStruct)
	fmt.Println("x: ", x)

	// 直接初始化
	y := demoStruct{
		Name:   "小明",
		Mobile: "13811111111",
		Gender: 0,
		code:   "abc",
	}
	fmt.Println("y: ", y)
}
