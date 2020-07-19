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

	rpc := RPCClient{}
	call := &Call{}
	call.Request = msg
	call.Address = "127.0.0.1:22221"


	rpc.Send(context.Background(), call)
	fmt.Println(call.Response)
	fmt.Println(call.Response.(map[string]interface{})["MethodName"])
}