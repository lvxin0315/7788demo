package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

func main() {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)
	for i := 0; i < 100; i++ {
		nc.Publish("foo", []byte(fmt.Sprintf("hi i = %d", i)))
		time.Sleep(time.Second)
	}
}
