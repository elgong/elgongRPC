package server

// server 接口
type Server interface {
	Server() error
	Register(rcvr interface{}, metaData map[string]string) error
	Close()
}
