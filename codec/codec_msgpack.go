package codec

import (
	"github.com/vmihailenco/msgpack"
	. "github.com/elgong/elgongRPC/plugin_centre"
)

// MsgPackCodec 注册进管理中心
func init(){
	msgpackCodec := MsgPackCodec{CodecType, "msgpackCodec"}
	PluginCenter.Register(msgpackCodec.Type, msgpackCodec.Name, msgpackCodec)
}

type MsgPackCodec struct{
	Type PluginType
	Name PluginName
}

func (c *MsgPackCodec) Encode(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (c *MsgPackCodec) Decode(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}