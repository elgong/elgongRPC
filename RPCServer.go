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
	// 1. 创建服务端
	server := server.NewRPCServer()

	// 2. 注册服务
	server.Register("MyService", new(MyService))

	// 3. 启动服务
	server.Server()

}
