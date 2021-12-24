package redisDB

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var(
	RedisPool *redis.Pool
	redisHost="127.0.0.1:6379"
)
func newRedisPool()*redis.Pool{
	return &redis.Pool{
		MaxIdle:50,         //最大空闲连接数
		MaxActive:30,        //允许分配最大连接数
		IdleTimeout:300*time.Second,    //连接时间限制
		Dial: func() (redis.Conn,  error) {    //创建连接的函数
			c,err := redis.Dial("tcp",redisHost ,redis.DialDatabase(1))
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			return c,nil
		},
	}
}



func init(){
	RedisPool = newRedisPool()
}