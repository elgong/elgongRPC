package main

import (
	"fmt"

	"github.com/elgong/elgongRPC/common"
	_ "github.com/elgong/elgongRPC/soa/register/redisPlugin"
)

func main() {

	//// 先注册
	//registerr := PluginCenter.Get(register.RegisterType, "redisRegisterPlugin").(register.Register)
	//
	//registerr.Register(context.Background(), "hehe", "1234.45.44.44")
	//time.Sleep(1000000000)
	//discoveyPlugin := PluginCenter.Get(discovey.DiscoveyType, "redisDiscoveyPlugin").(discovey.Discovey)
	//
	//ret := discoveyPlugin.Get("hehe")
	//
	//fmt.Println(ret)

	fmt.Println(common.GetClientIp())
}
