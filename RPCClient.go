package main

// demo
import (
	"context"
	"fmt"

	"github.com/elgong/elgongRPC/message"

	. "github.com/elgong/elgongRPC/client"
	_ "github.com/elgong/elgongRPC/soa/register/redisPlugin"
)

func main() {

	// 1. 创建客户端
	rpc := NewRpcClient()

	// 封装消息请求体
	msg := message.NewMessage()
	msg.SeqID = 11111
	msg.MethodName = "method"
	msg.Body = map[string]interface{}{"first": int64(1), "second": int64(2)}
	msg.ServiceName = "MyService"
	msg.MethodName = "Add"

	rsp := message.NewMessage()

	for i := 0; i < 10; i++ {
		// rpc 调用
		rpc.Call(context.Background(), msg, &rsp)

		fmt.Println("call return res: ", rsp.Body["res"])
	}

	// 阻塞
	select {}
}
