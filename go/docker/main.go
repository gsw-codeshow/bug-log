package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/", sayHello)
	log.Println("服务启动成功 监听端口：8000")
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if nil != err {
		log.Fatal("服务启动失败", err)
	}
	return
}
