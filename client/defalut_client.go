package client

// 客户端实现
import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/elgong/elgongRPC/config"

	"github.com/elgong/elgongRPC/soa/loadbalance"

	"github.com/elgong/elgongRPC/soa/discovey"

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

	select {
	case <-ctx.Done():
		fmt.Println("停止了...")
		return
	default:
	}
	// 对req 的封装，封装一个call，方便后续实现异步时使用
	call := &Call{}
	call.Request = reqBody
	call.ServiceName = call.Request.(message.DefalutMsg).ServiceName
	call.MethodName = call.Request.(message.DefalutMsg).MethodName

	// 服务发现插件
	discoveyPlugin := PluginCenter.Get("discovey", PluginName(config.DefalutGlobalConfig.DiscoveyPlugin)).(discovey.Discovey)

	// 获得可用的服务ip列表
	serviceList := discoveyPlugin.Get(call.ServiceName)

	// 负载均衡插件
	selector := PluginCenter.Get("selector", PluginName(config.DefalutGlobalConfig.SelectorPlugin)).(loadbalance.Selector)

	// 经过负载均衡选择的ip地址
	call.Address = selector.Select(serviceList)

	// 发送消息
	err := r.Send(ctx, call)

	// 发送失败，ip异常，上报错误并且移除ip
	if err != nil {

		discoveyPlugin.ReportAndRemove(call.ServiceName, call.Address)
		return
	}

	// 解析服务的返回值
	*rspBody = *call.Response.(*message.DefalutMsg)

}

func (r RPCClient) Send(ctx context.Context, call *Call) error {
	select {
	case <-ctx.Done():
		log.Println("停止了...")
		return nil
	default:
	}
	request := call.Request

	// 从插件管理中心 获取protocol 协议的编解码器
	proto := PluginCenter.Get("protocol", PluginName(config.DefalutGlobalConfig.ProtocolPlugin)).(protocol.Protocol)

	// 编码要发送的数据
	requestByte := proto.EncodeMessage(request)

	// 获取连接池插件，从连接池拿到 conn连接
	conn, err := PluginCenter.Get("connPool", PluginName(config.DefalutGlobalConfig.ConnPlugin)).(*DefaultPools).GetConn(ctx, call.Address)

	// 防止空的
	defer func(err error) {
		if err != nil {
			return
		}
		conn.PutBack()
	}(err)

	if err != nil {
		log.Println("网络异常:" + err.Error())
		// r.pendingCalls.Delete(seq)
		call.Error = err
		// call.done()

		return errors.New("conn err")
	}

	_, err = conn.Send(requestByte)

	fmt.Println("发送完成")

	if err != nil {
		log.Println("client write error:" + err.Error())
		// r.pendingCalls.Delete(seq)
		call.Error = err
		call.done()
		return err
	}

	log.Println("read response....")
	// 读取响应
	msg, err := proto.DecodeMessage(conn.Conn)

	if err != nil {
		fmt.Println("data resolve err")
	}
	call.Response = msg

	return nil
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
