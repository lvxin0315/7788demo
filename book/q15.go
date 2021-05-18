package main

import "fmt"

// 人类性别的枚举
const (
	Unknown = 0 // 未知
	Male    = 1 // 男性
	Female  = 2 // 女性
)

func main() {
	people := Female

	if people == Male {
		fmt.Println("我去男厕所")
	} else {
		fmt.Println("我去女厕所")
	}
}
