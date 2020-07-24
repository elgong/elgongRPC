package conn_pool

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// 等着补单测,时间不够,先临时凑合测试.... 复杂写就是testify, 简单写就是convey
func TestUnitConnDeque(t *testing.T) {

	Convey("ConnDeque 双向链表测试", t, func() {
		// connDeq 创建一个底层的双向链表
		connDeq := connDeque{}

		Convey("异常逻辑：空值弹出", func() {
			// popBottom
			for i := 0; i <= 3; i++ {
				conn := connDeq.popBottom()
				So(conn, ShouldResemble, (*connInPool)(nil))
			}

			// pop
			for i := 0; i <= 3; i++ {
				conn := connDeq.pop()
				So(conn, ShouldResemble, (*connInPool)(nil))
			}

			// poll
			for i := 0; i <= 3; i++ {
				conn := connDeq.poll()
				So(conn, ShouldResemble, (*connInPool)(nil))
			}
		})

		//Convey("异常逻辑：空值弹出", func() {
		//
		//})
	})
}
