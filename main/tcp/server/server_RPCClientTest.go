package main

import (
	"fmt"

	. "github.com/elgong/elgongRPC/plugin_centre"
	. "github.com/elgong/elgongRPC/protocol"

	//"io"
	//"log"
	"net"
	//"strings"
	"time"
)

func main() {
	//1.建立监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:22221")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	fmt.Println("listen Start...:")
	codec := PluginCenter.Get("protocol", "defaultProtocol").(Protocol)

	for {

		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)

		}
		for {
			//2.接收客户端的链接

			msg, _ := codec.DecodeMessage(conn)

			fmt.Println("-------------")
			fmt.Println("server shoudao :  ", msg)

			m := msg.(*DefalutMsg)
			m.ServiceName = "shoudesdsa dsad sads dasdsad "

			//byt := PluginCenter.Get("protocol", "defaultProtocol").(protocol.Protocol).EncodeMessage(m)
			// codec := PluginCenter.Get("protocol", "defaultProtocol").(protocol.Protocol)
			byt := codec.EncodeMessage(m)

			_, err := conn.Write(byt)
			if err != nil {
				fmt.Println(err)
				break
			}
			//for i := 0; i < len(byt); {
			//	n, err := conn.Write(byt[i:])
			//	if err != nil {
			//		fmt.Println(err)
			//	}
			//	i += n
			//	fmt.Println(i)
			//}

		}
	}

}

//处理请求，类型就是net.Conn
func process(conn net.Conn) {

	//处理结束后关闭链接
	// defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v", err)
			break
		}
		fmt.Printf(time.Now().String()+"recv from client, content:%v\n", string(buf[:n]))
	}

}
