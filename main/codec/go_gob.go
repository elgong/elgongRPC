// gob 使用
package main

import (
"bytes"
"encoding/gob"
"fmt"
	"io"
)

type MsgData struct {
	X, Y, Z int
	Name string
}


func main() {

	// network 网络传递的数据载体
	var network bytes.Buffer

	// 编码数据
	sendMsg := MsgData{1, 2, 3, "elgongRPC-V1.0"}

	err := senMsg(sendMsg, &network)
	if err!=nil {
		fmt.Println("编码错误")
		return
	}

	// 解码数据
	reciveMsg := MsgData{}
	err = revMsg(reciveMsg, &network)
	if err!=nil {
		fmt.Println("解码错误")
		return
	}
}

func senMsg( msg MsgData, w io.Writer)error {
	fmt.Print("开始执行编码（发送端）")

	enc := gob.NewEncoder(w)

	fmt.Println("原始数据：", msg)
	err := enc.Encode(&msg)
	fmt.Println("传递的编码数据为：", w)
	return  err
}
func revMsg(rev MsgData, r io.Reader)error {
	dec:=gob.NewDecoder(r)
	err:= dec.Decode(&rev) //传递参数必须为 地址
	fmt.Println("解码之后的数据为：",rev)
	return err
}
