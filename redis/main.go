package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Subscribe() {
	defer wg.Done()

	conn, err := Redis.Dail("tcp", "123.57.191.4:6379")
	if err != nil {
		fmt.Println("Redis.Dail err:", err)
		return
	}
	defer conn.Close()

	//实例化一个redigo中的PubSubConn结构体
	c := Redis.PubSubConn{
		Conn: conn,
	}
	//订阅一个频道
	err = c.Subscribe("channel1")
	if err != nil {
		fmt.Println("Subscribe err:", err)
		return
	}
	//接收redis服务器返回的数据
	//循环接收
	for {
		//Receive()返回的是空接口interface{}的类型,所以需要断言
		switch v := c.Receive().(type) {
		//Redis.Message结构体
		//type Message struct {
		//	Channel string
		//	Pattern string
		//	Data    []byte
		//}
		case Redis.Message:
			fmt.Printf("channel:%s,\tdata:%S\n", v.Channel, string(v.Data))

		//订阅或者取消订阅
		case Redis.Subscription:
			fmt.Printf("Channel:%s\tCount:%d \tKind:%s\n", n.Channel, n.Count, n.Kind)
		}
	}
}

func main() {
	//重新创建一个链接
	//不能在订阅的连接上
	conn, err := Redis.Dail("tcp", "123.57.191.4:6379")
	if err != nil {
		fmt.Println("Redis.Dail err:", err)
		return
	}
	defer conn.Close()
	//起一个协程
	wg.Add(1)
	go Subscribe()
	conn.Do("publish", "channel1", "hello")
	wg.Wait()

}
