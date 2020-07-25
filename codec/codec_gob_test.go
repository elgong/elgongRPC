package codec

import (
	"fmt"
	"testing"

	. "github.com/elgong/elgongRPC/plugin_centre"
	"github.com/smartystreets/assertions"
)

func TestUnitGobCodec(t *testing.T) {
	// 字符串验证
	// 注册进中心
	gobCodec := GobCodec{CodecType, "gobCodec"}
	PluginCenter.Register(gobCodec.Type, gobCodec.Name, &gobCodec)

	data := "hello world !!!" // 待传输的数据
	// 编码
	byt, _ := PluginCenter.Get(CodecType, "gobCodec").(Codec).Encode(data)

	// 解码
	var res string
	PluginCenter.Get(CodecType, "gobCodec").(Codec).Decode(byt, &res)

	// 断言
	assertions.ShouldEqual(res, data)

	// 结构体验证
	dataStruct := DefalutMsg{}
	dataStruct.Body = map[string]string{
		"name": "elgong",
	}
	dataStruct.MethodName = "哈哈哈"
	dataStruct.Body = map[string]interface{}{"name": "elgong"}

	// 编码
	byt2, _ := PluginCenter.Get(CodecType, "gobCodec").(Codec).Encode(dataStruct)

	// 解码
	var res2 DefalutMsg
	PluginCenter.Get(CodecType, "gobCodec").(Codec).Decode(byt2, &res2)

	fmt.Println(res2)
	assertions.ShouldEqual(res2, dataStruct)
}
