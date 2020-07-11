package main

import (
	"fmt"
	"net"
)

func main() {
	//1.建立一个链接（Dial拨号）
	conn, err := net.Dial("tcp", "127.0.0.1:8084")

	if err != nil{
		fmt.Println("errrrrr")

		return
	}

	conn.Write([]byte("123344"))
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


	fmt.Println("结束")
}