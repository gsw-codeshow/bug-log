package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool

type TransferConn struct {
	redis.Conn
	TransferConn redis.Conn
}

func (t TransferConn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	fmt.Println(" --- --- 1 --- --- ")
	fmt.Println(commandName, " --- --- ", args)
	reply, err = t.Conn.Do(commandName, args...)
	// 由于Key需要变化，所以不能单独只修改这个Do的部分，要根据取出之后的值来判断。
	t.TransferConn.Do("HMSET", args, reply)

	fmt.Println(reply)
	return
}

func DialTransferConn() redis.Conn {
	conn, _ := redis.Dial("tcp", "127.0.0.1:6379")
	password := "ginibong"
	if password != "" {
		if _, authErr := conn.Do("AUTH", password); authErr != nil {
			panic(authErr)
		}
	}
	conn.Do("SELECT", 9)
	return conn
}

// 初始化redis链接池
func RedisInit() {
	redisPool = &redis.Pool{
		MaxIdle:     100,
		MaxActive:   1024,
		IdleTimeout: time.Second * 10,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			// 123.57.191.4:6379
			var transfer TransferConn
			conn, err := redis.Dial("tcp", "127.0.0.1:1921")
			if err != nil {
				return conn, err
			}
			password := "123456"
			if password != "" {
				if _, authErr := conn.Do("AUTH", password); authErr != nil {
					return nil, fmt.Errorf("redis auth password error: %s", authErr)
				}
			}
			conn.Do("SELECT", 9)
			transfer.Conn = DialTransferConn()
			transfer.TransferConn = conn
			return transfer, nil
		},
	}
}

func main() {
	RedisInit()
	con := redisPool.Get()
	reply, err := con.Do("HGETALL", "UNIT:6744205645341265920:6747440582701789184")
	fmt.Println(reply)
	fmt.Println(err)
	fmt.Println(con)
	return
}
