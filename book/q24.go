package main

import "fmt"

func main() {
	i := 0
	// 我是一个标记
Label1:
	fmt.Println("我是Label1")
	for true {
		i++
		if i == 5 {
			goto Label1
		}
		if i > 10 {
			goto Label2
		}
		// 默认情况下，我们执行这里
		fmt.Println("默认的输出当前i的值：", i)
	}
	// 我是另一个标记
Label2:
	fmt.Println("我是Label2")
}
