// @Title 服务定义
// @Author  elgong 2020.7.18
// @Update  elgong 2020.7.21
package server

// Server 接口
type Server interface {
	Server() error
	Register(rcvr interface{}, metaData map[string]string) error
	Close()
}
