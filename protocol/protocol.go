package protocol
// 消息协议插件  接口
import (
	"io"
)

const ProtocolType = "protocol"

// 消息定义接口
type Protocol interface {

	DecodeMessage(r io.Reader) (interface{}, error)
	EncodeMessage(message interface{}) []byte
}
