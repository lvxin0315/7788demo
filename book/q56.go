package main

import (
	"errors"
	"fmt"
)

func main() {
	// 返回值是error的内置interface类型
	err := errors.New("自定义错误")
	fmt.Println("err: ", err)
}
