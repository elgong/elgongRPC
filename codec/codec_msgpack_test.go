package codec

import (
	. "github.com/elgong/elgongRPC/plugin_centre"
	"github.com/smartystreets/assertions"
	"testing"
)

func TestUnitMsgPackCodec(t *testing.T) {
	// 注册进中心
	msgpackCodec := MsgPackCodec{CodecType, "msgpackCodec"}
	PluginCenter.Register(msgpackCodec.Type, msgpackCodec.Name, msgpackCodec)

	data := "hello world !!!"  // 待传输的数据
	// 编码
	byt, _ := PluginCenter.Get(CodecType, "msgpackCodec").(Codec).Encode(data)

	// 解码
	var res string
	PluginCenter.Get(CodecType, "msgpackCodec").(Codec).Decode(byt, &res)

	assertions.ShouldEqual(res, data)
}
