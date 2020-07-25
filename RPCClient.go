package main

import (
	"context"
	"fmt"

	"github.com/elgong/elgongRPC/message"

	. "github.com/elgong/elgongRPC/client"
)

func main() {

	// 1. 创建客户端
	rpc := NewRpcClient()

	// 封装消息请求体
	msg := message.NewMessage()
	msg.SeqID = 11111
	msg.MethodName = "method"
	msg.Body = map[string]interface{}{"name": "elgong"}
	msg.ServiceName = "MyService"
	msg.MethodName = "PrintA"

	rsp := message.NewMessage()
	fmt.Println("发送前", msg)
	// rpc 调用
	rpc.Call(context.Background(), msg, &rsp)

	fmt.Println("shoudao ", rsp)

	// 阻塞
	select {}
}
