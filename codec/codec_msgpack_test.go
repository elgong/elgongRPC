package codec

import (
	. "github.com/elgong/elgongRPC/plugin_centre"
	"github.com/smartystreets/assertions"
	"testing"
)

// TestUnitMsgPackCodec  msgpack 插件的单元测试
// 连同插件管理中心一起测了。。。先偷懒一下，未来再修改
func TestUnitMsgPackCodec(t *testing.T) {

	// 注册进管理中心
	msgpackCodec := MsgPackCodec{CodecType, "msgpackCodec"}
	PluginCenter.Register(msgpackCodec.Type, msgpackCodec.Name, &msgpackCodec)

	// 字符串验证
	data := "hello world !!!"  // 待传输的数据
	// 编码
	byt, _ := PluginCenter.Get(CodecType, "msgpackCodec").(Codec).Encode(data)

	// 解码
	var res string
	PluginCenter.Get(CodecType, "msgpackCodec").(Codec).Decode(byt, &res)

	// 断言
	assertions.ShouldEqual(data, res)

	// 结构体验证
	dataStruct := DefalutMsg{}
	dataStruct.Body = map[string]string{
		"name":"elgong",
	}
	dataStruct.MethodName = "哈哈哈"
	// 编码
	byt2, _ := PluginCenter.Get(CodecType, "msgpackCodec").(Codec).Encode(dataStruct)

	dataStructRes := DefalutMsg{}
	// 解码
	PluginCenter.Get(CodecType, "msgpackCodec").(Codec).Decode(byt2, &dataStructRes)

	// 断言
	assertions.ShouldEqual(dataStructRes, dataStruct)

}

type DefaultMsgHeader struct {
	SeqID           uint64      // 序号, 用来唯一标识请求或响应
	IsRequest   	bool        // 消息类型，用来标识一个消息是请求还是响应
	SerializeType 	string      //序列化类型，用来标识消息体采用的编码方式
	IsException     bool        //状态类型，用来标识一个请求是正常还是异常
	ServiceName   string        //服务名
	MethodName    string        //方法名
	Error         string        //方法调用发生的异常
	OtherData      map[string]interface{} //其他想要携带的数据
}

// DefalutMsg 默认消息
type DefalutMsg struct {
	// 消息头部
	DefaultMsgHeader

	// 消息体
	Body interface{}
}
