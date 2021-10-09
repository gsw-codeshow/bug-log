package main

import (
	"fmt"
	"reflect"
)

func main() {
	r := reflect.ValueOf(1)
	if !r.IsValid() {
		fmt.Println(" --- --- 1 --- ---")
	}
	fmt.Println(r.IsValid())
	if r.IsZero() {
		fmt.Println(" --- --- 2 --- ---")
	}
	return
}
