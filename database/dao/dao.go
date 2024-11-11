package dao

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/tidwall/gjson"
	"gone/database/access"
	"gone/pkg/utils"
	"time"
)

// Rds 封装redis及其操作
type Rds struct {
	rdc redis.Conn // redis链接
	key string     // 要存入的key
}

// NewRedis 获取新的redis链接
func NewRedis(key string) *Rds {
	return &Rds{
		rdc: access.RedisPool.Get(), // 从 redis 连接池中获取一个连接
		key: key,                    // 要存入的key
	}
}

/*
	redis的key-value组合式存储，有过期时间
*/

// SetRedisKey 向redis存值
func (r *Rds) SetRedisKey(value interface{}, ex int) {
	// 将value转为[]byte
	vb, _ := json.Marshal(value)
	// 将[]byte存入
	_, err := r.rdc.Do("SET", r.key, vb, "EX", ex)
	if err != nil {
		fmt.Println("向redis存值失败: ", err)
	}
	// fmt.Println("向redis存值成功: ", reply, "key: ", r.key)
	return
}

// SearchAndGet 向redis查值，若查到则返回
func (r *Rds) SearchAndGet() (num int, value interface{}) {
	num = r.IsRedisKey()
	if num != 0 {
		value = r.GetRedisKey()
		return 1, value
	}
	return 0, value
}

// IsRedisKey 向redis查值
func (r *Rds) IsRedisKey() (haveField int) {
	reply, err := redis.Int(r.rdc.Do("EXISTS", r.key))
	if err != nil {
		fmt.Println("向redis查值失败: ", err)
	}
	return reply
}

// GetRedisKey 从redis获取值
func (r *Rds) GetRedisKey() (value interface{}) {
	// 查询key的[]byte值
	vb, err := redis.Bytes(r.rdc.Do("GET", r.key))
	if err != nil {
		fmt.Println("从redis获取值失败: ", err)
	}
	// 将[]byte转为interface{}类型
	value = gjson.Parse(string(vb)).Value()
	return value
}

// DelRedisKey 在redis删除值
func (r *Rds) DelRedisKey() (reply int) {
	reply, err := redis.Int(r.rdc.Do("DEL", r.key))
	if err != nil {
		fmt.Println("从redis删除值失败: ", err)
	}
	return reply
}

// CloseRedis 关闭线程池
func (r *Rds) CloseRedis() {
	err := r.rdc.Close() // 关闭连接池
	if err != nil {
		fmt.Println(err)
	}
}

/*
	以下是hash，由于hash的field，没有过期时间
*/

// 设置时间，用以计算
func ptime(i int) time.Duration {
	str := fmt.Sprintf("%vs", i)
	tm, _ := time.ParseDuration(str)
	return tm
}

// HashSetRedisKey 向redis存储hash值，增加时间参数，存储时填入
func (r *Rds) HashSetRedisKey(field, value interface{}, tin int) {
	newValue := map[string]interface{}{
		"field": value,                      // 参数
		"while": time.Now().Add(ptime(tin)), // 时间
	}
	// 将value转为[]byte
	vByte, _ := json.Marshal(newValue)
	// 将[]byte存入
	reply, err := r.rdc.Do("HSet", r.key, field, vByte)
	if err != nil {
		fmt.Println("向redis存hash值失败: ", err)
	} else {
		fmt.Printf("向redis存hash值成功: %v key: %v field: %v\n", reply, r.key, field)
	}
	return
}

// HashSearchAndGet 向redis查询hash值
func (r *Rds) HashSearchAndGet(field interface{}) (num int, value interface{}) {
	// 查询
	s := utils.AnyToString(field)
	numb := r.hashIsRedisKey(s)
	if numb != 0 {
		// 获取
		value = r.hashGetRedisKey(s)
		if value == nil {
			// 如果为空则直接返回
			return 0, nil
		}
		valueStr := utils.AnyToString(value)
		// 判断
		wt := gjson.Get(valueStr, "while").Time()
		if wt.Sub(time.Now()) <= 0 {
			// 没时间的话，可以删除了
			_ = r.hashDelRedisKey(s)
			return 0, nil
		}
		nv := gjson.Get(valueStr, "field")
		return 1, nv
	}
	return 0, nil
}

// 向redis查询hash值
func (r *Rds) hashIsRedisKey(field interface{}) (reply int) {
	reply, err := redis.Int(r.rdc.Do("HExists", r.key, field))
	if err != nil {
		fmt.Println("向redis查hash值失败: ", err)
	}
	return reply
}

// 从redis获取hash值
func (r *Rds) hashGetRedisKey(field interface{}) (value interface{}) {
	// 查询key的[]byte值
	vb, err := redis.Bytes(r.rdc.Do("HGet", r.key, field))
	if err != nil {
		fmt.Println("从redis获取hash值失败: ", err)
	}
	// 将[]byte转为interface{}类型
	value = gjson.Parse(string(vb)).Map()
	return value
}

// HashDelRedisKey 在redis删除hash值
func (r *Rds) hashDelRedisKey(field interface{}) (reply int) {
	reply, err := redis.Int(r.rdc.Do("HDel", r.key, field))
	if err != nil {
		fmt.Println("从redis删除hash值失败: ", err)
	}
	return reply
}
