// @Title 编解码插件gob
// @Author  elgong 2020.7.18
// @Update  elgong 2020.7.18
package codec

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/elgong/elgongRPC/message"

	. "github.com/elgong/elgongRPC/plugin_centre"
)

// GobCodec 注册进插件管理中心
func init() {
	gobCodec := GobCodec{CodecType, "gobCodec"}
	PluginCenter.Register(gobCodec.Type, gobCodec.Name, &gobCodec)
}

type GobCodec struct {
	Type PluginType
	Name PluginName
}

// body 类型单独注册
type interfaceType = map[string]interface{}

func (g GobCodec) Encode(value interface{}) ([]byte, error) {
	var buf bytes.Buffer
	_, ok := value.(message.DefalutMsg)
	if ok {
		fmt.Println("注册")
		gob.Register(interfaceType{})
	}

	err := gob.NewEncoder(&buf).Encode(value)
	fmt.Println(err)
	return buf.Bytes(), err
}

func (g GobCodec) Decode(data []byte, value interface{}) error {

	gob.Register(interfaceType{})
	buf := bytes.NewBuffer(data)
	err := gob.NewDecoder(buf).Decode(value)
	return err
}

///////////////////// 为什么实现接口， 不能使用 g *GobCodec, 提示未实现？？？
