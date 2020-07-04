// content of client.go
package main

import (
"fmt"
"log"
"net/rpc"
)

// Args 服务接受的参数
type Args struct {
	A, B int
}

// Quotient 服务返回的参数
type Quotient struct {
	Quo, Rem int
}


const(
	tcpAddr = "127.0.0.1:8080"
	httpAddr = "127.0.0.1:8081"
)

func main() {

	// 连接到一个HTTP RPC服务
	client, err := rpc.DialHTTP("tcp", httpAddr)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	// 入参
	args := &Args{7, 8}
	// 返回参数
	var reply int

	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	// 连接到一个RPC服务
	clientTCP, err := rpc.Dial("tcp", tcpAddr)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Asynchronous call
	// 返回参数
	quotient := new(Quotient)

	divCall := clientTCP.Go("Arith.Divide", args, &quotient, nil)

	// will be equal to divCall
	replyCall := <-divCall.Done
	if replyCall.Error != nil {
		fmt.Println(replyCall.Error)
	} else {
		fmt.Printf("Arith: %d/%d=%d...%d\n", args.A, args.B, quotient.Quo, quotient.Rem)
	}
}