package main

import (
	"fmt"

	"github.com/elgong/elgongRPC/message"
	. "github.com/elgong/elgongRPC/plugin_centre"
	"github.com/elgong/elgongRPC/protocol"

	"net"
)

func main() {
	//1.建立一个链接（Dial拨号）
	conn, err := net.Dial("tcp", "127.0.0.1:22222")

	if err != nil {
		fmt.Println("errrrrr")

		return
	}

	//if err != nil {
	//	fmt.Printf("dial failed, err:%v\n", err)
	//	return
	//}
	//
	//fmt.Println("Conn Established...:")
	//conn.LocalAddr()
	//
	////读入输入的信息
	//reader := bufio.NewReader(os.Stdin)
	//for {
	//	data, err := reader.ReadString('\n')
	//	if err != nil {
	//		fmt.Printf("read from console failed, err:%v\n", err)
	//		break
	//	}
	//
	//	data = strings.TrimSpace(data)
	//	//传输数据到服务端
	//	_, err = conn.Write([]byte(data))
	//	if err != nil {
	//		fmt.Printf("write failed, err:%v\n", err)
	//		break
	//	}
	//}

	msg := message.NewMessage()
	msg.SeqID = 11111
	msg.MethodName = "method"
	msg.Body = map[string]string{"name": "elgong"}
	fmt.Println(msg)

	byte := PluginCenter.Get("protocol", "defaultProtocol").(protocol.Protocol).EncodeMessage(msg)
	conn.Write(byte)
	// fmt.Println(byte)

	fmt.Println("结束")
}
