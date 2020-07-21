package protocol

import . "github.com/elgong/elgongRPC/plugin_centre"

// 消息体头部  可定制

func init() {

	PluginCenter.Register(MsgType, "defaultMsg", &DefalutMsg{})
}

// NewMessage 创建MSG 对象
func NewMessage() DefalutMsg {
	return DefalutMsg{}
}

// DefaultMsgHeader  默认消息头
type DefaultMsgHeader struct {
	SeqID         uint64 // 序号, 用来唯一标识请求或响应
	IsRequest     bool   // 消息类型，用来标识一个消息是请求还是响应
	SerializeType string //序列化类型，用来标识消息体采用的编码方式
	IsException   bool   //状态类型，用来标识一个请求是正常还是异常
	ServiceName   string //服务名
	MethodName    string //方法名
	Error         string //方法调用发生的异常
}

// DefalutMsg 默认消息
type DefalutMsg struct {
	// 消息头部
	DefaultMsgHeader

	// 消息体
	Body interface{}
}
