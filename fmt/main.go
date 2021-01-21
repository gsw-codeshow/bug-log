package main

import (
	"fmt"
	"math/rand"
	"time"
)

var randIndex = rand.New(rand.NewSource(time.Now().Unix()))

func RandInt(max int) int {
	return randIndex.Intn(max)
}

func main() {
	str := fmt.Sprintf("%%%d", 1)
	fmt.Println(str)
	for i := 0; i < 1000; i++ {
		fmt.Println(RandInt(2))
	}
	return
}
