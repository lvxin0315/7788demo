package main

import (
	"errors"
	"fmt"
)

func main() {
	err := funcDemo9()
	if err != nil {
		fmt.Println("报错了")
	}
	// 另一种方式
	if err := funcDemo9(); err != nil {
		fmt.Println("又报错了")
	}
}

func funcDemo9() error {
	return errors.New("err")
}
