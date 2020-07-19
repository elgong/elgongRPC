// 2020.7.8
package client

// 客户端接口
import (
	"context"
)

// Client 接口
type Client interface {
	Call(ctx context.Context, reqBody interface{}, rspBody interface{})
	Close() error
	IsShutDown() bool
}
