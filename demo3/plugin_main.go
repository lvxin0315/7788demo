package main

import (
	"fmt"
	"plugin"
)

func main() {
	p1()
	p2()
	p3()
	p4()
}

//函数1，使用一个无参数函数
func p1() {
	//加载动态库
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}
	//在库中查找 函数1
	symbol1, err := p.Lookup("PDemo1")
	if err != nil {
		panic(err)
	}
	symbol1.(func())()
}

//函数2，使用一个有参数函数
func p2() {
	//加载动态库
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}
	//在库中查找 函数2
	symbol2, err := p.Lookup("PDemo2")
	if err != nil {
		panic(err)
	}
	symbol2.(func(str string))("test2")
}

//变量3，库中的int类型变量
func p3() {
	//加载动态库
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}
	//在库中查找 变量3
	symbol3, err := p.Lookup("PDemoVar3")
	if err != nil {
		panic(err)
	}
	//打印看看效果
	fmt.Println(*symbol3.(*int))
	//修改PDemoVar3
	var test2 = 10
	*symbol3.(*int) = test2
	//再打印看看效果
	fmt.Println(*symbol3.(*int))
}

//函数3，修改库中int变量，然后调用函数3输出变量
func p4() {
	//加载动态库
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}
	//在库中查找 变量3
	symbol3, err := p.Lookup("PDemoVar3")
	if err != nil {
		panic(err)
	}
	//在库中查找 函数3
	symbol4, err := p.Lookup("PDemo3")
	if err != nil {
		panic(err)
	}
	//修改PDemoVar3
	var test2 = 10
	*symbol3.(*int) = test2
	//再打印看看效果
	fmt.Println(*symbol3.(*int))
	//调用函数3试试效果
	symbol4.(func())()
}
