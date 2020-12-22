// gin-vue/consumer.go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

var start int64

func doConsumerTask(i int) {
	// 1. 创建消费者
	config := nsq.NewConfig()
	q, errNewCsmr := nsq.NewConsumer("one-test", "one-test", config)
	if errNewCsmr != nil {
		fmt.Printf("fail to new consumer!, topic=%s, channel=%s", "one-test")
	}

	// 2. 添加处理消息的方法
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		// log.Printf("message: %v", string(message.Body))
		// log.Printf("%d", i)
		if "0" == string(message.Body) {
			start = time.Now().Unix()
		}
		if "4999" == string(message.Body) {
			log.Printf("time: %d", time.Now().Unix()-start)
		}
		message.Finish()
		return nil
	}))

	// 3. 通过http请求来发现nsqd生产者和配置的topic（推荐使用这种方式）
	lookupAddr := []string{
		"127.0.0.1:4161",
	}
	err := q.ConnectToNSQLookupds(lookupAddr)
	if err != nil {
		log.Panic("[ConnectToNSQLookupds] Could not find nsqd!")
	}

	// 4. 接收消费者停止通知
	<-q.StopChan

	// 5. 获取统计结果
	stats := q.Stats()
	fmt.Sprintf("message received %d, finished %d, requeued:%s, connections:%s",
		stats.MessagesReceived, stats.MessagesFinished, stats.MessagesRequeued, stats.Connections)
}

func main() {
	for i := 0; i < 5; i++ {
		go doConsumerTask(i)
	}
	select {}
}
