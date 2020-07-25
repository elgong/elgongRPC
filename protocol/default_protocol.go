package protocol

// 消息协议插件
import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/elgong/elgongRPC/config"

	"github.com/elgong/elgongRPC/message"

	"github.com/elgong/elgongRPC/codec"
	. "github.com/elgong/elgongRPC/plugin_centre"
)

// magicNum 魔数
var magicNum = []byte{0xaa, 0xbb}

// rpcVersion rpc 版本
var rpcVersion = "v1.0"

// 注册插件到 插件管理中心
func init() {
	deafultProtocol := DefaultProtocol{ProtocolType, "defaultProtocol"}
	PluginCenter.Register(deafultProtocol.Type, deafultProtocol.Name, &deafultProtocol)
}

// DefaultProtocol 默认协议
type DefaultProtocol struct {
	Type PluginType
	Name PluginName
}

// EncodeMessage 编码到二进制
func (d DefaultProtocol) EncodeMessage(msgP interface{}) []byte {

	codec := PluginCenter.Get(codec.CodecType, PluginName(config.DefalutGlobalConfig.CodecPlugin)).(codec.Codec)

	if codec == nil {
		panic("插件未注册")
	}

	msg, err := codec.Encode(msgP.(message.DefalutMsg))
	if err != nil {
		panic("编码信息失败")
	}

	rpcVersionByte, _ := codec.Encode(rpcVersion)

	// 别忘记加自身
	headLen := len(magicNum) + len(rpcVersionByte) + 2 + 4
	headLenBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(headLenBytes, uint16(headLen))

	// 别忘记加自身
	totalLen := headLen + len(msg)
	totalLenBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(totalLenBytes, uint32(totalLen))

	// data 二进制数据
	data := make([]byte, totalLen)
	start := 0

	copyFullWithOffset(data, magicNum, &start)

	copyFullWithOffset(data, rpcVersionByte, &start)

	copyFullWithOffset(data, totalLenBytes, &start)

	copyFullWithOffset(data, headLenBytes, &start)

	copyFullWithOffset(data, msg, &start)

	//fmt.Println("encode")
	//fmt.Println("totle len:", totalLen)
	//fmt.Println("head len", headLen)
	//fmt.Println("data len", len(msg))
	return data

}

func (d DefaultProtocol) DecodeMessage(r io.Reader) (interface{}, error) {

	var err error
	msg := &message.DefalutMsg{} // PluginCenter.Get("msg", "defaultMsg")
	codec := PluginCenter.Get(codec.CodecType, PluginName(config.DefalutGlobalConfig.CodecPlugin)).(codec.Codec)

	if codec == nil {
		panic("插件未注册")
	}

	// 解析魔数
	first2Byte := make([]byte, 2)
	_, err = io.ReadFull(r, first2Byte)
	if err != nil {
		fmt.Println("解析魔数时网络错误")
		return msg, err
	}

	if first2Byte[0] != magicNum[0] || first2Byte[1] != magicNum[1] {
		fmt.Println("魔数错误")
		return msg, errors.New("魔数错误")
	}

	// 解析 rpc 版本
	rpcVersionByte, _ := codec.Encode(rpcVersion)

	rpcV := make([]byte, len(rpcVersionByte))
	_, err = io.ReadFull(r, rpcV)
	if err != nil {
		fmt.Println("解析RPC version时, 出问题了")
		return msg, err
	}

	// 解析总长度
	totalLenBytes := make([]byte, 4)
	_, err = io.ReadFull(r, totalLenBytes)
	if err != nil {
		fmt.Println("解析 totalLenBytes 时, 出问题了")
		return msg, err
	}
	totalLen := int(binary.BigEndian.Uint32(totalLenBytes))

	if totalLen < 4 {
		err = errors.New("invalid total length")
		return msg, err
	}

	// 解析头长度
	headLenBytes := make([]byte, 2)
	_, err = io.ReadFull(r, headLenBytes)
	if err != nil {
		fmt.Println("解析 headLenBytes 时, 出问题了")
		return msg, err
	}

	headLen := int(binary.BigEndian.Uint16(headLenBytes))

	data := make([]byte, totalLen-headLen)
	_, err = io.ReadFull(r, data)
	if err != nil {
		fmt.Println(err)
		return msg, err
	}

	codec.Decode(data, msg)
	//fmt.Println("decode")
	//fmt.Println("totle len:", totalLen)
	//fmt.Println("head len", headLen)
	//fmt.Println("data len", len(data))
	return msg, nil
}

func copyFullWithOffset(dst []byte, src []byte, start *int) {
	copy(dst[*start:*start+len(src)], src)
	*start = *start + len(src)
}
