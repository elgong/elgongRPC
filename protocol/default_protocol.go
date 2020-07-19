package protocol
// 消息协议插件
import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/elgong/elgongRPC/codec"
	. "github.com/elgong/elgongRPC/plugin_centre"
	"io"
)

var magicNum = []byte{0xaa, 0xbb}
var rpcVersion = "v1.0"

// 注册插件到 插件管理中心
func init(){

	deafultProtocol := DefaultProtocol{ProtocolType, "defaultProtocol"}
	PluginCenter.Register(deafultProtocol.Type, deafultProtocol.Name, &deafultProtocol)
}

// DefaultProtocol 默认协议
type DefaultProtocol struct {
	Type PluginType
	Name PluginName
}

// EncodeMessage 编码到二进制
func (d DefaultProtocol) EncodeMessage(message interface{}) []byte{
	//
	codec := PluginCenter.Get(codec.CodecType, "msgpackCodec").(codec.Codec)

	if codec == nil {
		panic("插件未注册")
	}

	msg, err := codec.Encode(message)
	if err != nil{
		panic("编码信息失败")
	}

	rpcVersionByte, _ := codec.Encode(rpcVersion)

	// 别忘记加自身
	headerLen := len(magicNum) + len(rpcVersionByte) + 2
	headerLenBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(headerLenBytes, uint16(headerLen))

	// 别忘记加自身
	totalLen := headerLen + len(msg) + 4
	totalLenBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(totalLenBytes, uint32(totalLen))

	// data 二进制数据
	data := make([]byte, totalLen)
	start := 0
	copyFullWithOffset(data, magicNum, &start)
	copyFullWithOffset(data, rpcVersionByte, &start)
	copyFullWithOffset(data, totalLenBytes, &start)
	copyFullWithOffset(data, headerLenBytes, &start)

	copyFullWithOffset(data, msg, &start)
	fmt.Println("encode")
	fmt.Println("总",totalLen)
	fmt.Println("头", headerLen)
	fmt.Println("data", len(data))
	return data

}

func (d DefaultProtocol) DecodeMessage(r io.Reader) (interface{}, error) {

	var err error
	msg := &DefalutMsg{}// PluginCenter.Get("msg", "defaultMsg")
	codec := PluginCenter.Get(codec.CodecType, "msgpackCodec").(codec.Codec)


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
		return msg, err
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

	data := make([]byte, totalLen - headLen - 4)
	_, err = io.ReadFull(r, data)

	if err != nil {
		fmt.Println(err)
	}



	codec.Decode(data, msg)
	fmt.Println("decode")
	fmt.Println("总",totalLen)
	fmt.Println("头", headLen)
	fmt.Println("data", len(data))
	return msg, nil
}

func copyFullWithOffset(dst []byte, src []byte, start *int) {
	copy(dst[*start:*start+len(src)], src)
	*start = *start + len(src)
}
