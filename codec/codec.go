package codec
// 编解码插件  编解码统一接口

const CodecType = "codec"

// Codec 编解码接口
type Codec interface {
	Encode(value interface{}) ([]byte, error)
	Decode(data []byte, value interface{}) error
}
