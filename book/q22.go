package main

import "fmt"

func main() {
	i := 0
	// 只要i小于10
	for i < 10 {
		fmt.Println("现在的i是：", i)
		// 自增没忘记吧
		i++
	}
}
