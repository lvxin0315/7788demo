package main

import (
	"fmt"
	"github.com/lvxin0315/7788demo/book/demo"
)

func main() {
	// 初始化
	x := new(demo.DataDemo)
	code := x.InitCode()
	fmt.Println("code: ", code)
}
