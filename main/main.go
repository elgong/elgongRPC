package main

import (
	"context"
	"fmt"
	"github.com/elgong/elgongRPC/interceptor"
)
func call (ctx context.Context, serviceName string, methodName string, reqBody interface{}, rspBody interface{}) error {

	fmt.Println("函数调用了")
	return nil
}


func in (ctx context.Context, serviceName string, methodName string, reqBody interface{}, rspBody interface{}, callFunc interceptor.CallFunc) error {

	fmt.Println("之前执行哦")
	call(ctx, serviceName, methodName, reqBody, rspBody)
	return nil
}
func main() {

	 interceptor.Interceptorss.Register(in)
	 //interceptor.Interceptorss.Register(in)

	fmt.Println(len(*interceptor.Interceptorss))

	_ = interceptor.Interceptorss.GetInterceptors(context.Background(), "123", "123", "123", "123", call)(context.Background(),"123","123","123","123")
	
}
