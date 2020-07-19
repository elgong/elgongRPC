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
	msg.Body = map[string]string{"name":"elgong"}
	msg.ServiceName = "service-2"
	msg.MethodName = "printB"

	rpc := RPCClient{}

	msg2 := protocol.NewMessage()
	rpc.Call(context.Background(), msg, msg2)

	fmt.Println("shoudao  ", msg2)
	// fmt.Println(msg.Response.(map[string]interface{})["MethodName"])

	select {

	}
}