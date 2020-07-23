package main

import (
	"context"
	"fmt"

	. "github.com/elgong/elgongRPC/client"
	"github.com/elgong/elgongRPC/protocol"
)

func main() {

	// 1. 创建客户端
	rpc := NewRpcClient()

	// 封装消息请求体
	msg := protocol.NewMessage()
	msg.SeqID = 11111
	msg.MethodName = "method"
	msg.Body = map[string]string{"name": "elgong"}
	msg.ServiceName = "MyService"
	msg.MethodName = "PrintA"

	rsp := protocol.NewMessage()

	// rpc 调用
	rpc.Call(context.Background(), msg, &rsp)

	fmt.Println("shoudao ", rsp)

	// 阻塞
	select {}
}
