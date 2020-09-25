package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

//生产者
func main() {
	//第一步 连接rabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//MQ使用的channel
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	//我们循环起来，让发送不停止
	for i := 0; i < 100; i++ {
		//发送内容
		body := fmt.Sprintf("当前时间戳:%d", time.Now().Unix())
		err = ch.ExchangeDeclare(
			"logs_topic", // name
			"topic",      // type
			true,         // durable
			false,        // auto-deleted
			false,        // internal
			false,        // no-wait
			nil,          // arguments
		)
		if err != nil {
			panic(err)
		}
		err = ch.Publish("demoServ", //「交换机」名称
			"all", //消息主题
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("发送第%d次， 内容是 %s", i+1, body))
		//sleep一下，太快看不出来效果
		time.Sleep(500 * time.Millisecond)
	}
	//发完了100次
	fmt.Println("发完了100次")
}
