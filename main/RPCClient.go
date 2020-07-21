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
	msg.ServiceName = "MyService"
	msg.MethodName = "PrintA"

	rpc := NewRpcClient()

	msg2 := protocol.NewMessage()
	msg2.IsRequest = true
	rpc.Call(context.Background(), msg, &msg2)

	fmt.Println("shoudao ", msg2)

	select {}
}
