package main

import (
	"fmt"

	"github.com/elgong/elgongRPC/protocol"
	"github.com/elgong/elgongRPC/server"
)

type MyService struct {
}

// 这里入参和出餐的断言有问题哦
func (m *MyService) PrintA(msg *protocol.DefalutMsg, msg2 *protocol.DefalutMsg) {
	fmt.Println("调用了...printA")
	fmt.Println("request:  ", msg.Body.(map[string]interface{})["name"])
	msg2.Body.(map[string]string)["name"] = "eeeee1111"
}
func main() {
	server := server.NewRPCServer()
	server.Register("MyService", new(MyService))
	server.Server()

}
