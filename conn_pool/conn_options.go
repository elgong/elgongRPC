// @Title 连接池的参数设置
// @Description  函数选项模式
// @Author  elgong 2020.7.
// @Update  elgong 2020.7.
package conn_pool

import (
	"fmt"

	"github.com/elgong/elgongRPC/config"
)

var defaultConnOptions ConnOptions

// 从全局配置导入到该配置
func init() {

	fmt.Println(config.DefalutGlobalConfig)
	// defaultConnOptions 默认参数
	defaultConnOptions = ConnOptions{
		initialCap: config.DefalutGlobalConfig.Conn.InitialCap,
		maxCap:     config.DefalutGlobalConfig.Conn.MaxCap,  // 0 默认关闭
		maxIdle:    config.DefalutGlobalConfig.Conn.MaxIdle, // 0 默认关闭
		//idletime:  1,
		//maxLifetime: 2,
		failReconnect:       config.DefalutGlobalConfig.Conn.FailReconnect,
		failReconnectSecond: config.DefalutGlobalConfig.Conn.FailReconnectSecond,
		failReconnectTime:   config.DefalutGlobalConfig.Conn.FailReconnectTime,
		isTickerOpen:        config.DefalutGlobalConfig.Conn.IsTickerOpen,
		tickerTime:          config.DefalutGlobalConfig.Conn.TickerTime,
	}
}

// ConnOptions 连接池参数结构体
type ConnOptions struct {
	initialCap int
	maxCap     int
	maxIdle    int // 最大空闲时间  s
	//idletime    time.Duration
	//maxLifetime time.Duration

	// 超时重连
	failReconnect       bool // 是否掉线重连
	failReconnectSecond int  // 重连等待时间
	failReconnectTime   int  // 重连次数

	isTickerOpen bool // 定时任务是否开启
	tickerTime   int  // 定时秒数
}

//
type ModifyConnOption func(opt *ConnOptions)

// With*** 传入新参数，
func WithInitCap(initialCap int) ModifyConnOption {

	// opt *Options 传入待修改的参数指针
	return func(opt *ConnOptions) {
		opt.initialCap = initialCap
	}
}

func WithMaxCap(maxCap int) ModifyConnOption {

	// opt *Options 传入待修改的参数指针
	return func(opt *ConnOptions) {
		opt.maxCap = maxCap
	}
}

// 其他待补充
