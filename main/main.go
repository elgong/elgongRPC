package main

import (
	_ "github.com/elgong/elgongRPC/codec"
	"github.com/elgong/elgongRPC/message"
	"github.com/elgong/elgongRPC/server"
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

	msg := message.NewMessage()
	msg.SeqID = 11111
	msg.MethodName = "method"
	msg.Body = map[string]string{"name": "elgong"}
	msg.ServiceName = "service-1"
	msg.MethodName = "printA"

	rpc := server.NewRPCServer()

	rpc.Register("MyService", new(MyService))
	rpc.Invoke("MyService", "PrintA", &msg, &msg)

}
