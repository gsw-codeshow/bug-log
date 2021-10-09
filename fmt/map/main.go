package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func main() {
	onlineNum := make(map[string]int)
	lock := new(sync.RWMutex)
	g := sync.WaitGroup{}
	g.Add(1)
	go func() {
		defer g.Done()
		for i := 0; i < 100; i++ {
			select {
			case <-time.Tick(time.Second * 2):
				c := make(chan int, 10)
				name := RandString(3)
				for {
					lock.Lock()
					_, ok := onlineNum[name]
					lock.Unlock()
					if ok {
						name = RandString(3)
					} else {
						break
					}

				}
				go func() {
					lock.Lock()
					num := onlineNum[name]
					lock.Unlock()
					for {
						select {
						case <-c:
							lock.Lock()
							onlineNum[name] = num + 1
							num++
							lock.Unlock()
						}
					}
				}()

				go func() {
					tick := time.Tick(time.Second * 2)
					for {
						select {
						case <-tick:
							c <- r.Intn(10)
						}
					}
				}()

			}
		}
	}()
	go func() {
		tick := time.Tick(time.Second * 4)
		for {
			select {
			case <-tick:
				lock.Lock()
				fmt.Println("--- num ---")
				for name, num := range onlineNum {
					fmt.Println(name, "--- online num: ", num)
				}
				lock.Unlock()
			}
		}
	}()
	g.Wait()
	time.Sleep(1000 * time.Second)
	return
}
