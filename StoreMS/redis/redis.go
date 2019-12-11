package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strings"
)

var redisPool *redis.Pool

//建立redis连接
func GetRedisConn() redis.Conn {
	if redisPool == nil {
		redisPool = &redis.Pool{
			MaxActive:   100,
			MaxIdle:     10,
			IdleTimeout: 10,
			Dial: func() (redis.Conn, error) {
				conn, e := redis.Dial("tcp", "localhost:6379")
				conn.Do("AUTH", "123456")
				return conn, e
			},
		}
	}
	return redisPool.Get()
}

//关闭redis连接池
func CloseRedisConn(conn redis.Conn) {
	conn.Close()
}

//
func DoRedisCommand(cmd string, resultType string) (interface{}, error) {
	fmt.Println("DoRedisCommand...")
	conn := GetRedisConn()
	//
	strs := strings.Split(cmd, " ")
	args := make([]interface{}, 0)
	for _, arg := range strs[1:] {
		args = append(args, arg)
	}
	//
	reply, e := conn.Do(strs[0], args...)
	if reply != nil {
		switch resultType {
		case "string":
			return redis.String(reply, e)
		case "int":
			return redis.Int(reply, e)
		case "strings":
			return redis.Strings(reply, e)
			//default:
			//	return redis.String(reply, e)
		}
	}
	CloseRedisConn(conn)
	return reply, e
}
