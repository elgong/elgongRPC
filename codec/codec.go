package codec
// 类型定义

const CodecType = "codec"

// Codec 编解码接口
type Codec interface {
	Encode(value interface{}) ([]byte, error)
	Decode(data []byte, value interface{}) error
}
