package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func sig_handler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	<-ch
	fmt.Println("我不想被关闭")
	go sig_handler()
}

func main() {
	go sig_handler()
	time.Sleep(5 * time.Second)
	go func() { panic("测试堵塞") }()
	select {}
}
