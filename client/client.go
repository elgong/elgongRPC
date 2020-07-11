// 2020.7.8
package client

import "context"

// 接口
type Client interface {
	Call(ctx context.Context, serviceName string, method string, reqBody interface{}, rspBody interface{}) error
}


type client struct {


}

// Call 调用服务后端
func (c client) Call(ctx context.Context, serviceName string, method string, reqBody interface{}, rspBody interface{}) error{

	// 调用包装器
}


func (c client) send(ctx context.Context, serviceName string, method string, reqBody interface{}, rspBody interface{}){

	// codec 编码

	// transport 发送


}