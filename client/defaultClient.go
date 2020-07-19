package client

import (
	"context"
	"fmt"
	"github.com/elgong/elgongRPC/codec"
	. "github.com/elgong/elgongRPC/conn_pool"
	. "github.com/elgong/elgongRPC/plugin_centre"
	"github.com/elgong/elgongRPC/protocol"
	"log"
	"sync"
)

type RPCClient struct {
	seq          uint64
	codec        codec.Codec
	pendingCalls sync.Map
	mutex        sync.Mutex
	shutdown     bool
}

func (r *RPCClient) IsShutDown() bool {
	return r.shutdown
}

// Call 调用服务后端
func (r RPCClient) Call(ctx context.Context, serviceName string, method string, reqBody interface{}, rspBody interface{}) error{

	// 调用包装器
	return nil
}


func (r RPCClient) Send(ctx context.Context, call *Call){

	//seq := r.seq
	//r.pendingCalls.Store(seq, call)

	request := call.Request

	proto := PluginCenter.Get("protocol", "defaultProtocol").(protocol.Protocol)

	requestByte := proto.EncodeMessage(request)

	conn, err := PluginCenter.Get("connPool", "defaultConnPool").(*DefaultPools).GetConn(call.Address)

	if err != nil {
		log.Println("网络异常:" + err.Error())
		// r.pendingCalls.Delete(seq)
		call.Error = err
		call.done()
	}

	_, err  = conn.Send(requestByte)
	fmt.Println("发送完成")

	if err != nil {
		log.Println("client write error:" + err.Error())
		// r.pendingCalls.Delete(seq)
		call.Error = err
		call.done()
		return
	}

	fmt.Println("读取响应")
	// 读取响应
	msg, err := proto.DecodeMessage(conn.Conn)

	fmt.Println("读取完成")

	if err != nil {
		fmt.Println("数据解析错误")
	}

	call.Response = msg
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

type Call struct {
	Address string
	ServiceName   string
	MethodName string      // 服务名.方法名
	Request          interface{} // 参数
	Response         interface{} // 返回值（指针类型）
	Error         error       // 错误信息
	Done          chan *Call  // 在调用结束时激活
}

func (c *Call) done() {
	c.Done <- c
}
