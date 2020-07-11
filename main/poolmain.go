package main

import (
	"fmt"
	"github.com/elgong/elgongRPC/conn_pool"
	"time"
)

func main(){
	for {
		conn, err := conn_pool.Pool.GetConn("127.0.0.1:8088")
		conn2, err2 := conn_pool.Pool.GetConn("127.0.0.1:8004")

		if err != nil {
			fmt.Println("errrrrr")

			return
		}

		if err2 != nil {
			fmt.Println("2222222err")
			return
		}
		

		conn.Conn.Write([]byte("conn1:------1234"))
		conn2.Conn.Write([]byte("conn2------4567"))
		conn.PutBack()
		conn2.PutBack()

		time.Sleep(300)
		fmt.Println("完成")
	}


}
