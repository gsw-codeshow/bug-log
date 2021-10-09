package main

import "fmt"

func DeferDebug() {
	fmt.Println("---DeferDebug---")
	return
}

func main() {
	defer DeferDebug()
	for i := 0; i < 100; i++ {
		goto Error
	}

Error:
	fmt.Println("Error")
	return
}
