package conn_pool

import (
	"fmt"
	"testing"
)

// 等着补单测,时间不够,先临时凑合测试....
func TestUnitConnDeque(t *testing.T) {
	deque := connDeque{}

	conn := deque.popBottom()

	if conn == nil {
		fmt.Println("nil")
	}

	deque.push(connInPool{address: "1"})
	deque.push(connInPool{address: "2"})
	deque.push(connInPool{address: "3"})

	for i := 0; i < 3; i++ {
		conn := deque.popBottom()

		fmt.Println(conn.address)
	}
}
