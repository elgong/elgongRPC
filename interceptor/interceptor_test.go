package interceptor

import (
	"context"
	"fmt"
	"testing"
)

func call (ctx context.Context, serviceName string, methodName string, reqBody interface{}, rspBody interface{}) error {

	fmt.Println("函数调用了")
	return nil
}


func in (ctx context.Context, serviceName string, methodName string, reqBody interface{}, rspBody interface{}, callFunc CallFunc) error {

	fmt.Println("函数之前拦截啦")
	call(ctx, serviceName, methodName, reqBody, rspBody)
	return nil
}
func TestUnitInterceptors(t *testing.T) {

	Interceptorss.Register(in)
	// Interceptorss.Register(in)

	_ = Interceptorss.GetInterceptors(context.Background(), "123", "123", "123", "123", call)

}
