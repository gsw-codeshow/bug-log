package main

import "fmt"

func sayHi() {
	fmt.Println("hi")
	return
}

func sayPanic() {
	panic("panic")
}

// 线程中的panic，也会导致进程消失。是最高级别的警告。除非有recover来恢复上下文，最好不要。
func main() {
	for i := 0; i < 50; i++ {
		go sayHi()
	}
	for i := 0; i < 10; i++ {
		go sayPanic()
	}
	select {}
}
