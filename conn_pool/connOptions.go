package conn_pool

import (
	"time"
)

// ConnOptions 连接池参数结构体
type ConnOptions struct {
	initialCap  int
	maxCap      int
	maxIdle     int
	idletime    time.Duration
	maxLifetime time.Duration
}

// defaultConnOptions 默认参数
var defaultConnOptions = ConnOptions{
	initialCap:  10,
	maxCap:    10,
	maxIdle:   10,
	idletime:  1,
	maxLifetime: 2,
}

//
type ModifyConnOption func(opt *ConnOptions)

// With*** 传入新参数，
func WithInitCap(initialCap int) ModifyConnOption{

	// opt *Options 传入待修改的参数指针
	return func(opt *ConnOptions){
		opt.initialCap = initialCap
	}
}

func WithMaxCap(maxCap int) ModifyConnOption{

	// opt *Options 传入待修改的参数指针
	return func(opt *ConnOptions){
		opt.maxCap = maxCap
	}
}

// 其他待补充

