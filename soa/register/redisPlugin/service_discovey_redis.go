// @Title  redis 实现的服务发现
// @Author  elgong 2020.8.1
// @Update  elgong 2020.8.1
package redisPlugin

import (
	"fmt"

	. "github.com/elgong/elgongRPC/plugin_centre"
	"github.com/elgong/elgongRPC/soa/discovey"

	"github.com/garyburd/redigo/redis"
)

func init() {

	opt := redis.DialPassword(defaultRedisOptions.password)

	conn, err := redis.Dial("tcp", defaultRedisOptions.ip, opt)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		panic("Redis Connect to redis error")
	}
	redisDiscovey := RedisDiscovey{discovey.DiscoveyType, "redisDiscoveyPlugin", []string{}, conn}
	PluginCenter.Register(redisDiscovey.Type, PluginName(redisDiscovey.Name), &redisDiscovey)
}

// RedisDiscovey reids 服务发现插件
type RedisDiscovey struct {
	Type         PluginType
	Name         PluginName
	serviceCache []string   // 本地的可用服务缓存
	conn         redis.Conn // redis 链接
}

// Get 获取列表
func (r RedisDiscovey) Get(serviceName string) []string {

	// 如果本地缓存为空，则可能第一次，去取一下数据吧
	if len(r.serviceCache) == 0 {
		r.serviceCache = r.getFromRedis(serviceName)
	}
	return r.serviceCache
}

// getFromRedis 从redis 获取
func (r *RedisDiscovey) getFromRedis(serviceName string) []string {
	result, err := redis.Values(r.conn.Do("smembers", "namespace_"+serviceName))

	if err != nil {
		fmt.Println("服务注册失败")
		return nil
	}
	ret := []string{}
	redis.ScanSlice(result, &ret)
	return ret
}
