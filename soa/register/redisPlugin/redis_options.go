// @Title redis的参数设置
// @Description  函数选项模式
// @Author  elgong 2020.7.
// @Update  elgong 2020.7.
package redisPlugin

import "regexp"

var defaultRedisOptions RedisOptions

func init() {
	defaultRedisOptions = RedisOptions{
		ip:          "127.0.0.1:6379",
		serviceName: "test",
		password:    "1234",
	}
}

// ConnOptions 连接池参数结构体
type RedisOptions struct {
	// Ip 地址
	ip string
	// 服务名
	serviceName string
	// 密码
	password string
	// 上报时间间隔 1s
	time int
}

//
type ModifyRedisOptions func(opt *RedisOptions)

// With*** 传入新参数，
func WithIP(Ip string) ModifyRedisOptions {
	// 格式校验
	match, _ := regexp.MatchString("((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}", Ip)

	if !match {
		Ip = "127.0.0.1:6379"
	}
	// opt *Options 传入待修改的参数指针
	return func(opt *RedisOptions) {
		opt.ip = Ip
	}
}

func WithServiceName(serviceName string) ModifyConnOption {

	// opt *Options 传入待修改的参数指针
	return func(opt *RedisOptions) {
		opt.serviceName = serviceName
	}
}

func WithPassword(password string) ModifyConnOption {

	// opt *Options 传入待修改的参数指针
	return func(opt *RedisOptions) {
		opt.password = password
	}
}
