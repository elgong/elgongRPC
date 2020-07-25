// @Title server的参数设置
// @Description  函数选项模式
// @Author  elgong 2020.7.25
// @Update  elgong 2020.7.25
package server

import (
	"fmt"

	"github.com/elgong/elgongRPC/config"
)

var defaultServerOptions ServerOptions

// 从全局配置导入到该配置
func init() {

	// defaultServerOptions 默认参数
	defaultServerOptions = ServerOptions{
		Ip:   config.DefalutGlobalConfig.Server.Ip,
		Port: config.DefalutGlobalConfig.Server.Port,
	}
	fmt.Println("server参数解析成功")
}

// ConnOptions 连接池参数结构体
type ServerOptions struct {
	Ip   string
	Port string
}

//
type ModifyServerOptions func(opt *ServerOptions)

// With*** 传入新参数，
func WithIp(ip string) ModifyServerOptions {

	// opt *Options 传入待修改的参数指针
	return func(opt *ServerOptions) {
		opt.Ip = ip
	}
}

func WithPort(port string) ModifyServerOptions {

	// opt *Options 传入待修改的参数指针
	return func(opt *ServerOptions) {
		opt.Port = port
	}
}

// 其他待补充
