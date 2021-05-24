package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("开始")
	// 同时开始任务1和2
	go funcDemo10()
	funcDemo11()
	// 到此等待一下任务1完成
	time.Sleep(5 * time.Second)
}

// 任务1 输出1-100的数字
func funcDemo10() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}

// 任务2 输出666 100次
func funcDemo11() {
	for i := 1; i <= 100; i++ {
		fmt.Println("我是666")
	}
}
