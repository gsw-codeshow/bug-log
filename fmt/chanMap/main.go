package main

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var nameIndex int
var nameLock sync.RWMutex

func GetName() (name string) {
	nameLock.Lock()
	nameIndex++
	nameLock.Unlock()
	num := strconv.Itoa(nameIndex)
	name = "chan_" + num
	return
}

func main() {
	mapLock := sync.RWMutex{}
	chanMap := make(map[string]int)

	go func() {
		for {
			select {
			case <-time.Tick(time.Second * 1):
				name := GetName()
				go func() {
					for {
						select {
						case <-time.Tick(time.Second * 1):
							time.Sleep(time.Second * time.Duration(rand.Int()%10))
							mapLock.Lock()
							c := chanMap[name]
							chanMap[name] = c + 1
							mapLock.Unlock()
						}
					}
				}()

			}
		}
	}()
	go func() {
		for {
			select {
			case <-time.Tick(time.Second * 2):
				mapLock.Lock()
				m := chanMap
				println(" -------- ")
				for c, n := range m {
					println(c, "--", n)
				}
				println(" --------- ")
				mapLock.Unlock()
			}
		}
	}()
	select {}
	return
}
