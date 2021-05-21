package main

import (
	"fmt"
	"github.com/lvxin0315/7788demo/book/demo"
)

func main() {
	// 初始化
	x := new(demo.DataDemo)
	x.Name = "小明"
	fmt.Println("name: ", x.Name)

	// TODO 下面代码会报错
	fmt.Println("code: ", x.code)
}
