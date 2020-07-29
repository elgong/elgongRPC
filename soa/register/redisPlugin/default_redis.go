// @Title  redis 实现的服务注册
// @Author  elgong 2020.7.29
// @Update  elgong 2020.7.29
package redisPlugin

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/elgong/elgongRPC/soa/register"

	. "github.com/elgong/elgongRPC/plugin_centre"

	"github.com/garyburd/redigo/redis"
)

func init() {
	redis := RedisRegister{register.RegisterType, "redisPlugin"}

	PluginCenter.Register(redis.Type, PluginName(redis.Type), &redis)
}

// RedisRegister redis 注册插件
type RedisRegister struct {
	Type PluginType
	Name PluginName
}

// Register 实现注册
func (r *RedisRegister) Register(ctx context.Context, serviceName string, Ip string, opts ...ModifyRedisOptions) {

	// opt 默认参数
	var redisOpt = defaultRedisOptions
	// 根据传入参数调整
	for _, o := range opts {
		o(&redisOpt)
	}

	// 连接redis数据库,指定数据库的IP和端口
	opt := redis.DialPassword(redisOpt.password)
	conn, err := redis.Dial("tcp", redisOpt.ip, opt)
	defer conn.Close()

	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	// 定时时长
	ticker := time.NewTicker(time.Second * time.Duration(redisOpt.time))

	wait := sync.WaitGroup{}
	wait.Add(1)
	go func(ctx context.Context) {
		for _ = range ticker.C {

			select {
			case <-ctx.Done():
				fmt.Println("停止了...")
				return
			default:
			}

			_, err = conn.Do("SET", "mykey", "superRobot")
			if err != nil {
				fmt.Println("redis set failed:", err)
			}
		}
		wait.Done()
	}(ctx)

	wait.Wait()
	fmt.Println("上报结束")

}
