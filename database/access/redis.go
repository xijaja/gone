package access

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gomodule/redigo/redis"
	"gone/config"
)

var RedisPool *redis.Pool // 创建redis连接池

// 初始化赋值
func init() {
	RedisPool = redisPool()
}

// 创建 redis 线程池
func redisPool() *redis.Pool {
	return &redis.Pool{ // 实例化一个连接池
		MaxIdle:     16,       // 最初地连接数量
		MaxActive:   0,        // 连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配，最多为100万。
		IdleTimeout: 300,      // 连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial:        dialInit, // 要连接的redis数据库
	}
}

// 连接初始化
func dialInit() (redis.Conn, error) {
	// 连接 redis 服务器
	address := fmt.Sprintf("%s:%s", config.Config.Redis.Host, config.Config.Redis.Port)
	c, err := redis.Dial("tcp", address)
	if err != nil {
		log.Error(err.Error())
	}
	// 传入密码
	if _, err2 := c.Do("AUTH", config.Config.Redis.Pass); err2 != nil {
		log.Error(err2.Error())
	}
	// 选择数据库
	_, err3 := c.Do("SELECT", 0)
	if err3 != nil {
		log.Error("Redis 恐慌：", err3.Error())
	}
	return c, nil
}
