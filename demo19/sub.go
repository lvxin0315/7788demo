package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"runtime"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)

	nc.Subscribe("foo", func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
	})

	runtime.Goexit()
}
