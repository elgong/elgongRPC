package main

import (
	"fmt"
	"github.com/elgong/elgongRPC/conn_pool"
	"time"
)

func main(){
	for {
		conn, err := conn_pool.Pool.GetConn("127.0.0.1:11111")
		conn2, err2 := conn_pool.Pool.GetConn("127.0.0.1:22222")
		// conn.SetDead()
		if err != nil {
			fmt.Println("errrrrr")
		} else {
			n1, err11 := conn.Conn.Write([]byte("conn1:------1234"))
			if n1 == 0 || err11 != nil {
				conn.SetDead()
			}
			conn.PutBack()
		}

		if err2 != nil {
			fmt.Println("errrrrr")
		} else{
			n2, err22 := conn2.Conn.Write([]byte("conn2------4567"))

			if n2== 0 || err22 != nil {
				conn2.SetDead()
			}
			conn2.PutBack()
		}




		//conn.Conn.SetWriteDeadline(time.Now().Add(time.Microsecond * 10))
		//conn2.Conn.SetWriteDeadline(time.Now().Add(time.Microsecond * 10))




		time.Sleep(300000)
		// fmt.Println("完成")
	}


}
