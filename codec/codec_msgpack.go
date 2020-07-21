// @Title 编解码插件 msgpack
// @Author  elgong 2020.7.18
// @Update  elgong 2020.7.18
package codec

import (
	. "github.com/elgong/elgongRPC/plugin_centre"
	"github.com/vmihailenco/msgpack"
)

// MsgPackCodec 注册进管理中心
func init() {
	msgpackCodec := MsgPackCodec{CodecType, "msgpackCodec"}
	PluginCenter.Register(msgpackCodec.Type, msgpackCodec.Name, &msgpackCodec)
}

type MsgPackCodec struct {
	Type PluginType
	Name PluginName
}

func (c MsgPackCodec) Encode(value interface{}) ([]byte, error) {
	return msgpack.Marshal(value)
}

func (c MsgPackCodec) Decode(data []byte, value interface{}) error {
	return msgpack.Unmarshal(data, value)
}
