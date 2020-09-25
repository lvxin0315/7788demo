package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

//消费者
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

	//初始化几个消费者
	c1 := "name1"
	c2 := "name2"
	c3 := "name3"
	c4 := "name4"
	q1, _ := createQueue(ch, c1, "*")   //通配符，所有内容都接收
	q2, _ := createQueue(ch, c2, "all") //接收all的消息
	q3, _ := createQueue(ch, c3, "a")   //接收a的消息，应该什么都没有
	q4, _ := createQueue(ch, c4, "all") //接收all的消息
	//开始输出它们的内容
	go listenQueueBody(&q1, ch)
	go listenQueueBody(&q2, ch)
	go listenQueueBody(&q3, ch)
	go listenQueueBody(&q4, ch, " one")
	go listenQueueBody(&q4, ch, " two") //监听两次q4, 会分别收到消息

	select {}
}

//创建队列
func createQueue(ch *amqp.Channel, qName string, key string) (amqp.Queue, error) {
	//先删除
	_, _ = ch.QueueDelete(qName, false, false, true)
	//再创建
	q, err := ch.QueueDeclare(qName, true, false, false, true, nil)
	if err != nil {
		panic(err)
	}
	err = ch.QueueBind(qName, key, "demoServ", true, nil)
	if err != nil {
		panic(err)
	}
	return q, nil
}

//消费输出内容
func listenQueueBody(q *amqp.Queue, ch *amqp.Channel, args ...string) {
	delivery, err := ch.Consume(q.Name,
		"",
		true,  // auto ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil)   // args
	if err != nil {
		panic(err)
	}
	for d := range delivery {
		fmt.Println(fmt.Sprintf("我是%s, 内容是: %s", q.Name, string(d.Body)), args)
	}
}
