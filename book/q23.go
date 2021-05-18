package main

import "fmt"

func main() {
	i := 0
	for true {
		i++
		if i == 5 {
			fmt.Println("i 现在是5，我们continue，直接进入下一次循环")
			continue
		}
		if i > 10 {
			fmt.Println("i 现在大于10了，我们终止")
			break
		}
		// 默认情况下，我们执行这里
		fmt.Println("默认的输出当前i的值：", i)
	}
}
