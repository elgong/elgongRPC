package client

// 客户端实现
import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/elgong/elgongRPC/message"

	. "github.com/elgong/elgongRPC/conn_pool"
	. "github.com/elgong/elgongRPC/plugin_centre"
	"github.com/elgong/elgongRPC/protocol"
)

// NewRpcClient 创建rpc客户端
func NewRpcClient() *RPCClient {
	return &RPCClient{}
}

// RPCClient 客户端结构体
type RPCClient struct {
	seq          uint64 // message 编号
	pendingCalls sync.Map
	mutex        sync.Mutex
	shutdown     bool
}

// IsShutDown 判断是否断开
func (r RPCClient) IsShutDown() bool {
	return r.shutdown
}

// Call 调用服务后端
func (r RPCClient) Call(ctx context.Context, reqBody interface{}, rspBody *message.DefalutMsg) {
	// 对req 的封装，封装一个call，方便后续实现异步时使用
	call := &Call{}
	call.Request = reqBody

	// 服务发现， 发现服务地址，临时先直接赋值代替。
	// 获取IP 地址
	call.Address = "127.0.0.1:8999"

	r.Send(ctx, call)

	// 解析服务的返回值
	*rspBody = *call.Response.(*message.DefalutMsg)

}

func (r RPCClient) Send(ctx context.Context, call *Call) {

	//seq := r.seq
	//r.pendingCalls.Store(seq, call)

	request := call.Request

	// 从插件管理中心 获取protocol 协议的编解码器
	proto := PluginCenter.Get("protocol", "defaultProtocol").(protocol.Protocol)

	// 编码要发送的数据
	requestByte := proto.EncodeMessage(request)

	// 从连接池拿到 conn连接
	conn, err := PluginCenter.Get("connPool", "defaultConnPool").(*DefaultPools).GetConn(call.Address)
	defer conn.PutBack()
	if err != nil {
		log.Println("网络异常:" + err.Error())
		// r.pendingCalls.Delete(seq)
		call.Error = err
		call.done()
	}

	_, err = conn.Send(requestByte)

	fmt.Println("发送完成")

	if err != nil {
		log.Println("client write error:" + err.Error())
		// r.pendingCalls.Delete(seq)
		call.Error = err
		call.done()
		return
	}

	fmt.Println("read response....")
	// 读取响应
	msg, err := proto.DecodeMessage(conn.Conn)

	if err != nil {
		fmt.Println("data resolve err")
	}
	call.Response = msg
}

type Call struct {
	Address     string      // 地址  IP:PORT
	ServiceName string      // 要调用的服务名
	MethodName  string      // 方法名
	Request     interface{} // 传入参数
	Response    interface{} // 返回值（指针类型）
	Error       error       // 错误信息
	Done        chan *Call  // 在调用结束时使用，临时用不到
}

func (c *Call) done() {
	c.Done <- c
}

//func (r *RPCClient) Go(ctx context.Context, request interface{}, response interface{}, done chan *Call) *Call {
//	call := new(Call)
//	call.ServiceMethod = request.(*protocol.DefalutMsg).MethodName
//	call.ServiceName = request.(*protocol.DefalutMsg).ServiceName
//	call.Request = request
//	call.Response = response
//	call.address = "127.0.0.1:22221"
//
//	if done == nil {
//		done = make(chan *Call, 10) // buffered.
//	} else {
//		if cap(done) == 0 {
//			log.Panic("rpc: done channel is unbuffered")
//		}
//	}
//	call.Done = done
//
//	r.send(ctx, call)
//
//	return call
//}

//func (r RPCClient) input() {
//	var err error
//	var rsp interface{}
//	for err == nil {
//		rsp, err = PluginCenter.Get("protocol", "defaultProtocol").(protocol.Protocol).DecodeMessage(r.rwc)
//		if err != nil {
//			break
//		}
//		response := rsp.(*protocol.DefalutMsg)
//		seq := response.SeqID
//		callInterface, ok := r.pendingCalls.Load(seq)
//		if !ok {
//			//请求已经被清理掉了，可能是已经超时了
//			continue
//		}
//
//		call := callInterface.(*Call)
//
//		r.pendingCalls.Delete(seq)
//
//		switch {
//		case response.Error != "":
//			call.Error = errors.New("服务器响应错误")
//			call.done()
//		default:
//			call.Response = response
//			call.done()
//		}
//	}
//	log.Println("input error, closing client, error: " + err.Error())
//	// r.Close()
//}
