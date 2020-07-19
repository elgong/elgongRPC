package codec

import (
	. "github.com/elgong/elgongRPC/plugin_centre"
	"github.com/vmihailenco/msgpack"
)

// MsgPackCodec 注册进管理中心
func init(){
	msgpackCodec := MsgPackCodec{CodecType, "msgpackCodec"}
	PluginCenter.Register(msgpackCodec.Type, msgpackCodec.Name, &msgpackCodec)
}

type MsgPackCodec struct{
	Type PluginType
	Name PluginName
}

func (c MsgPackCodec) Encode(value interface{}) ([]byte, error) {
	return msgpack.Marshal(value)
}

func (c MsgPackCodec) Decode(data []byte, value interface{}) error {
	return msgpack.Unmarshal(data, value)
}
