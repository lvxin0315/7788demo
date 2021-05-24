package main

import (
	"fmt"
	"sync"
)

// 监工
var wg sync.WaitGroup

func main() {
	// 假设我们有10个worker在工作
	for i := 0; i < 10; i++ {
		// 监工开始让一个人干活
		wg.Add(1)
		go funcDemo12()
	}
	// 等大家都干完
	wg.Wait()

	fmt.Println("大家都干完了")
}

func funcDemo12() {
	// 开始干活了
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
	// 我干完了，告诉监工一下
	wg.Done()
}
