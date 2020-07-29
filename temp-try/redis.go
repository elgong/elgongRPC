package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//spring.redis.host=121.41.111.45
//spring.redis.password=Gelqq666%
//spring.redis.database=0
//spring.redis.port=6379
//spring.redis.timeout=5000
//spring.redis.jedis.pool.max-active=9
//spring.redis.jedis.pool.max-wait=3
//spring.redis.jedis.pool.min-idle=10
//spring.redis.jedis.pool.max-idle=10
func main() {
	// 连接redis数据库,指定数据库的IP和端口
	opt := redis.DialPassword("Gelqq666%")
	conn, err := redis.Dial("tcp", "121.41.111.45:6379", opt)

	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	} else {
		fmt.Println("Connect to redis ok.")
	}

	// 函数退出时关闭连接
	defer conn.Close()

	// 执行一个set插入
	_, err = conn.Do("SET", "mykey", "superRobot")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	// 读取指定set
	username, err := redis.String(conn.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}
