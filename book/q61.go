package main

import (
	"fmt"
	"time"
)

var event1 chan string
var event2 chan string
var event3 chan string

func main() {
	event1 = make(chan string)
	event2 = make(chan string)
	event3 = make(chan string)
	go funcDemo15()
	// 往channel中放入变量，就会执行对应的代码
	go func() {
		for true {
			event1 <- "a"
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for true {
			event2 <- "b"
			time.Sleep(2 * time.Second)
		}
	}()
	go func() {
		for true {
			event3 <- "c"
			time.Sleep(200 * time.Millisecond)
		}
	}()
	time.Sleep(5 * time.Second)
}

// 监听所有event
func funcDemo15() {
	for {
		select {
		case s := <-event1:
			fmt.Println("我是event1，s: ", s)
		case s := <-event2:
			fmt.Println("我是event2，s: ", s)
		case s := <-event3:
			fmt.Println("我是event3，s: ", s)
		}
	}
}
