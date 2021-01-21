package main

import (
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

func doConsumerTask(i int) {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("block", "block", config)
	if nil != err {
		log.Println("newConsumer error: ", err)
		return
	}

	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("message: %s", string(message.Body))
		go func() {
			log.Print("sleep")
			time.Sleep(5 * time.Second)
			log.Printf("--sleep--finish--")
		}()
		log.Println("---i---", i)
		time.Sleep(5 * time.Second)
		// message.Finish()
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
	for i := 0; i < 5; i++ {
		go doConsumerTask(i)
	}
	select {}
}
