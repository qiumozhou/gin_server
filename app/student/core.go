package student

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"server/component/redis"

)

func SetCache(data interface{})string{
	tData,_ := json.Marshal(data)
	conn := redisDB.RedisPool.Get()
	defer conn.Close()
	key := "student"
	res, _ := redis.String(conn.Do("SET", key, tData))
	return res
}


func GetCache(key string)(data string){
	conn := redisDB.RedisPool.Get()
	defer conn.Close()
	res, _ := redis.String(conn.Do("GET", key))
	return res
}







