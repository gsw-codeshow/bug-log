package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

func doConsumerTask() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("one-test", "one-test", config)
	if nil != err {
		log.Println("newConsumer error: ", err)
		return
	}

	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("message: %s", string(message.Body))
		message.Finish()
		return nil
	}))
	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if nil != err {
		log.Printf("connect lookupd err: %s", err)
	}
	<-consumer.StopChan
	return
}
func main() {
	doConsumerTask()
	return
}
