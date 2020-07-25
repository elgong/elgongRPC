package main

import (
	"fmt"

	"github.com/elgong/elgongRPC/message"

	"github.com/elgong/elgongRPC/server"
)

type MyService struct {
}

// 这里入参和出餐的断言有问题哦
func (m *MyService) PrintA(msg *message.DefalutMsg, msg2 *message.DefalutMsg) {
	fmt.Println("调用了...printA")

	fmt.Println("request:  ", msg.Body)
	msg2.Body["name"] = "eeeee1111"
}

// Add 计算两数之和
func (m *MyService) Add(msg *message.DefalutMsg, msg2 *message.DefalutMsg) {
	fmt.Println("调用了...Add")

	// 拿到要计算的值
	fisrt := msg.Body["first"].(int)
	second := msg.Body["second"].(int)

	res := fisrt + second
	msg2.Body["res"] = res
}

func main() {
	// 1. 创建服务端
	server := server.NewRPCServer()

	// 2. 注册服务
	server.Register("MyService", new(MyService))

	// 3. 启动服务
	server.Server()
}
