package server

// server 接口
type server interface {

	Server() error
}

type RPCServer struct {

}
