package main

import "fmt"

func main() {
	const (
		a = iota //	iota开始计数，此时是 0
		b        //	1
		c        //	2
		d = "hi" //	这里由于使用的"hi赋值",所以后续的e将不会继续使用iota的值，但是iota继续+1
		e        //	e会使用上次赋值的内容 "hi"   iota继续+1
		f = 100  //	iota继续+1
		g        //	100  iota继续+1
		h = iota //	此时的iota已经是7,恢复计数，所以h是7
		i        //	8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
