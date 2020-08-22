package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"math/rand"
	"sync"
	"time"
)

//我们做个比较耗时的function，这样能更好的看到执行情况
func demoFunc(i int) {
	//随机延时（单位：毫秒）
	s := rand.Intn(2000)
	time.Sleep(time.Duration(s) * time.Millisecond)
	fmt.Println(fmt.Sprintf("我是第%d个任务, 我的延时是%d秒 ... ok", i, s))
}

//加个轮询查看当前
func listenPoolRunningNum(p *ants.PoolWithFunc) {
	for true {
		time.Sleep(time.Second)
		fmt.Println(fmt.Sprintf("\n当前协程数量: %d\n", p.Running()))
	}
}

func main() {
	//release可以释放池
	defer ants.Release()
	//执行函数次数
	runTimes := 200
	//goroutine阻塞，否则我们就看不全结果了
	var wg sync.WaitGroup
	// 设置一下最大协程数，我们的函数需要在匿名函数中调用
	p, _ := ants.NewPoolWithFunc(16, func(i interface{}) {
		demoFunc(i.(int))
		wg.Done()
	})
	//挂个监听，看看效果
	go listenPoolRunningNum(p)
	defer p.Release()
	//开始疯狂执行
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		//这里就是在调用ants.Pool的匿名函数
		_ = p.Invoke(i)
	}
	wg.Wait()
}
