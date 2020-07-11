package main

import (
	"fmt"
	"net"
)

func main() {
	//1.建立监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:8004")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	fmt.Println("listen Start...:")

	for {
		//2.接收客户端的链接
		conn, err := listen.Accept()
		fmt.Println("11111111111111111111")
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		//3.开启一个Goroutine，处理链接
		go process(conn)
	}
}

func process(conn net.Conn) {

	//处理结束后关闭链接
	// defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v", err)
			break
		}
		fmt.Printf("recv from client, content:%v\n", string(buf[:n]))
	}

}




