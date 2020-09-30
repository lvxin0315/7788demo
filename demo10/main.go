package main

import (
	"fmt"
	"github.com/atotto/clipboard"
)

var content = ""

func main() {
	for {
		get()
	}
}

func get() {
	ctx, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	if content != ctx {
		content = ctx
		fmt.Println(ctx)
	}
}
