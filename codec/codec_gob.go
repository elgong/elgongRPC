// 2020.7.18
// v1 先实现 简单的，未来再优化
package codec

import (
	"bytes"
	"encoding/gob"
	. "github.com/elgong/elgongRPC/plugin_centre"
)

// GobCodec 注册进插件管理中心
func init(){
	gobCodec := GobCodec{CodecType, "gobCodec"}
	PluginCenter.Register(gobCodec.Type, gobCodec.Name, gobCodec)
}

type GobCodec struct {
	Type PluginType
	Name PluginName
}

func (g *GobCodec) Encode(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(v)
	return buf.Bytes(), err
}

func (g *GobCodec) Decode(data []byte, value interface{}) error {
	buf := bytes.NewBuffer(data)
	err := gob.NewDecoder(buf).Decode(value)
	return err
}



