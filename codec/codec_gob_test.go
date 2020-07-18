package codec

import (
	. "github.com/elgong/elgongRPC/plugin_centre"
	"github.com/smartystreets/assertions"
	"testing"
)

func TestUnitGobCodec(t *testing.T) {
	// 注册进中心
	gobCodec := GobCodec{CodecType, "gobCodec"}
	PluginCenter.Register(gobCodec.Type, gobCodec.Name, &gobCodec)

	data := "hello world !!!"  // 待传输的数据
	// 编码
	byt, _ := PluginCenter.Get(CodecType, "gobCodec").(Codec).Encode(data)

	// 解码
	var res string
	PluginCenter.Get(CodecType, "gobCodec").(Codec).Decode(byt, &res)

	assertions.ShouldEqual(res, data)
}
