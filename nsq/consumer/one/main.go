// gin-vue/consumer.go
package main

import (
	"fmt"
	"log"

	"github.com/nsqio/go-nsq"
)

var consumer *nsq.Consumer

func doConsumerTask() {
	// 1. 创建消费者
	config := nsq.NewConfig()
	consumer, _ = nsq.NewConsumer("one-test", "one-test", config)

	// 2. 添加处理消息的方法
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		//time.Sleep(time.Second * 10)
		log.Printf("message: %s", string(message.Body))
		message.Finish()
		return nil
	}))

	// 3. 通过http请求来发现nsqd生产者和配置的topic（推荐使用这种方式）
	lookupAddr := []string{
		"127.0.0.1:4161",
	}
	err := consumer.ConnectToNSQLookupds(lookupAddr)
	if err != nil {
		log.Panic("[ConnectToNSQLookupds] Could not find nsqd!")
	}

	// 4. 接收消费者停止通知
	<-consumer.StopChan

	// 5. 获取统计结果
	stats := consumer.Stats()
	fmt.Sprintf("message received %d, finished %d, requeued:%s, connections:%s",
		stats.MessagesReceived, stats.MessagesFinished, stats.MessagesRequeued, stats.Connections)
}

func main() {
	doConsumerTask()
}
