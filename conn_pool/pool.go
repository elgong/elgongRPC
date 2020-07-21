// @Title 连接池插件的接口
// @Description
// @Author  elgong 2020.7.
// @Update  elgong 2020.7.
package conn_pool

import "net"

const ConnType = "connPool"

// pool 连接池接口, 不同类型的连接池都要实现该接口
type pool interface {
	GetConn(address string) (net.Conn, error)
	Close() error
}
