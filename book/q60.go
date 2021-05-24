package main

import (
	"fmt"
	"time"
)

var ch chan int

func main() {
	// 初始化信道
	ch = make(chan int)
	// 开始任务
	go funcDemo13()
	go funcDemo13()
	go funcDemo14()

	time.Sleep(5 * time.Second)
}

// 通过协程将v放到信道
func funcDemo13() {
	for i := 0; i < 100; i++ {
		fmt.Println("往ch中放入：", i)
		ch <- i
	}
}

// 当有值放到信道用，就会执行
func funcDemo14() {
	for {
		v := <-ch
		fmt.Println("ch收到v：", v)
	}
}
