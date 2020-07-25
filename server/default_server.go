package server

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"sync"

	"github.com/elgong/elgongRPC/message"

	. "github.com/elgong/elgongRPC/protocol"

	. "github.com/elgong/elgongRPC/plugin_centre"
)

var typeOfError = reflect.TypeOf((*error)(nil)).Elem()
var typeOfContext = reflect.TypeOf((*context.Context)(nil)).Elem()

// NewRPCServer 创建并返回server
func NewRPCServer(opts ...ModifyServerOptions) *RPCServer {

	// opt 默认参数
	var opt = defaultServerOptions

	// 修改参数
	for _, o := range opts {
		o(&opt)
	}

	s := new(RPCServer)

	s.ip = opt.Ip
	s.port = opt.Port

	s.serviceMap = make(map[string]*Service)
	return s
}

// RPCServer 服务端
type RPCServer struct {
	// 服务的map 容器
	serviceMap map[string]*Service
	mutex      sync.Mutex
	// 是否关闭
	shutdown bool
	ip       string
	port     string
}

// Register 注册服务到容器中
func (r *RPCServer) Register(servcieName string, serviceImpl interface{}) {
	// 先创建一个空的服务
	service := &Service{}
	service.methodsMap = make(map[string]reflect.Method)
	service.rcvr = reflect.ValueOf(serviceImpl)
	service.typ = reflect.TypeOf(serviceImpl)
	service.serviceName = servcieName

	for i := 0; i < service.typ.NumMethod(); i++ {
		// 需要对方法检查吗
		method := service.typ.Method(i)
		methodName := method.Name
		service.methodsMap[methodName] = method
	}

	r.mutex.Lock()
	// 避免重复注册
	if _, OK := r.serviceMap[service.serviceName]; !OK {
		r.serviceMap[service.serviceName] = service
	}
	r.mutex.Unlock()
}

// Invoke 服务的方法调用
func (r *RPCServer) Invoke(serviceName string, methodName string, args interface{}, retArgs interface{}) {
	r.serviceMap[serviceName].invoke(methodName, args, retArgs)
}

func (r *RPCServer) Server() {
	//1.建立监听端口
	listen, err := net.Listen("tcp", r.ip+":"+r.port)
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	fmt.Println("server Start...:")

	for {
		//2.接收客户端的链接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)

		}

		r.handlerConn(&conn)

	}
}

func (r *RPCServer) handlerConn(conn *net.Conn) {

	codec := PluginCenter.Get("protocol", "defaultProtocol").(Protocol)
	for {

		fmt.Println("监听中")
		// 读 eof结束

		//2.接收客户端的链接
		msg, err := codec.DecodeMessage(*conn)
		if err != nil {
			fmt.Println("接受客户端连接失败")
			break
		}

		// fmt.Println("server shoudao :  ", msg)

		m := msg.(*message.DefalutMsg)

		// 必须一个完整的数据////////////////////////////////
		response := message.DefalutMsg{Body: map[string]interface{}{
			"name": "111111",
		}}
		// 调用
		r.Invoke(m.ServiceName, m.MethodName, m, &response)
		// fmt.Println("发送", response)
		byt := codec.EncodeMessage(response)
		_, err = (*conn).Write(byt)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
