package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	n, z := time.Now().Zone()
	fmt.Println(n, "---", z/60/60)
	atomic.CompareAndSwapInt32()
	return
}
