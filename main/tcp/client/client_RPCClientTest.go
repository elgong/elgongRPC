package main

import (
	"context"
	"fmt"

	. "github.com/elgong/elgongRPC/client"
	"github.com/elgong/elgongRPC/protocol"
)

func main() {

	msg := protocol.NewMessage()
	msg.SeqID = 11111
	msg.MethodName = "method"
	msg.Body = map[string]string{"name": "elgong"}
	msg.ServiceName = "service-1"
	msg.MethodName = "printA"

	rpc := NewRpcClient()

	msg2 := protocol.NewMessage()
	msg2.IsRequest = true
	rpc.Call(context.Background(), msg, &msg2)
	fmt.Println(msg2.ServiceName)

	msg2.IsRequest = true
	rpc.Call(context.Background(), msg, &msg2)
	fmt.Println(msg2.ServiceName)

	msg2.IsRequest = true
	rpc.Call(context.Background(), msg, &msg2)
	fmt.Println(msg2.ServiceName)

	msg2.IsRequest = true
	rpc.Call(context.Background(), msg, &msg2)
	fmt.Println(msg2.ServiceName)
	// fmt.Println(msg.Response.(map[string]interface{})["MethodName"])

	select {}
}
