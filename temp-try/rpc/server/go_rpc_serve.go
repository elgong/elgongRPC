// content of server_prototest.go
package main


import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Args 服务接受的参数
type Args struct {
	A, B int
}

// Quotient 服务返回的参数
type Quotient struct {
	Quo, Rem int
}

// Arith 服务
type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	fmt.Println("调用了Multiply。。。")
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	fmt.Println("调用了Divide。。。")
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func listenTCP(addr string) (net.Listener, string) {
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatalf("net.Listen tcp :0: %v", e)
	}
	return l, l.Addr().String()
}

func main() {
	const(
		tcpAddr = "127.0.0.1:8080"
		httpAddr = "127.0.0.1:8081"
	)

	// 注册服务
	rpc.Register(new(Arith))
	var l net.Listener

	// 监听TCP连接
	l, serverAddr := listenTCP(tcpAddr)
	log.Println("RPC server listening on", serverAddr)
	go rpc.Accept(l)

	// 监听HTTP连接
	rpc.HandleHTTP()
	l, serverAddr = listenTCP(httpAddr)
	log.Println("RPC server listening on", serverAddr)
	go http.Serve(l, nil)

	// 防止main 结束后关闭协程
	select{}
}