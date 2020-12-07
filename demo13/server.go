package main

import (
	"fmt"
	"net"
)

var clientMap map[string]net.Conn

//初始化
func init() {
	clientMap = make(map[string]net.Conn)
}

func main() {
	fmt.Println("start server...")
	listen, err := net.Listen("tcp", "0.0.0.0:50000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		//新的连接
		conn, err := listen.Accept() //监听是否有连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	fmt.Println("LocalAddr：", conn.LocalAddr())
	fmt.Println("RemoteAddr:", conn.RemoteAddr())
	for {
		buf := make([]byte, 512)
		//go func() {
		//	for  {
		//		conn.Write([]byte("hi"))
		//		time.Sleep(500 * time.Millisecond)
		//	}
		//}()
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			continue
		}
		fmt.Println("read: ", string(buf))

	}
}
