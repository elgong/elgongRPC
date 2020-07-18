// 2020.7.8
package client
// 客户端接口
import "context"

// 接口
type Client interface {
	Call(ctx context.Context, serviceName string, method string, reqBody interface{}, rspBody interface{}) error
	Close() error
	IsShutDown() bool
}
