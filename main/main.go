package main

import (
	"fmt"
	"time"

	_ "github.com/elgong/elgongRPC/codec"
)

//func call (ctx context.Context, serviceName string, methodName string, reqBody interface{}, rspBody interface{}) error {
//
//	fmt.Println("函数调用了")
//	return nil
//}
//
//
//func in (ctx context.Context, serviceName string, methodName string, reqBody interface{}, rspBody interface{}, callFunc interceptor.CallFunc) error {
//
//	fmt.Println("之前执行哦")
//	call(ctx, serviceName, methodName, reqBody, rspBody)
//	return nil
//}
func main() {

	// interceptor.Interceptorss.Register(in)
	// //interceptor.Interceptorss.Register(in)
	//
	//fmt.Println(len(*interceptor.Interceptorss))
	//
	//_ = interceptor.Interceptorss.GetInterceptors(context.Background(), "123", "123", "123", "123", call)(context.Background(),"123","123","123","123")
	//

	// plugin_centre.PluginCenter.Register()

	//// 插件中心
	//byt, _ := plugin_centre.PluginCenter.Get(codec.CodecType, "msgpackCodec").(codec.Codec).Encode("hello")
	//
	//fmt.Println(byt)
	//
	//var s string
	//
	//plugin_centre.PluginCenter.Get(codec.CodecType, "msgpackCodec").(codec.Codec).Decode(byt, &s)
	//
	//fmt.Println(s)

	s := time.Now()

	time.Sleep(5000000000)
	m := time.Now().Sub(s)
	fmt.Println(m)

	if m >= time.Second*3 {
		fmt.Println(33)
	}

}
