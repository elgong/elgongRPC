package conn_pool
// 连接池插件 接口

import "net"

const ConnType = "connPool"

// pool 连接池接口, 不同类型的连接池都要实现该接口
type pool interface {
	GetConn(address string) (net.Conn, error)
	Close() error
}
