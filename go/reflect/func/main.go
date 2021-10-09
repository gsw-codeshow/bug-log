package main

import (
	"fmt"
	"reflect"
)

type st struct {
}

func (this *st) Echo() {
	fmt.Println("echo(----)")
	return
}

func main(){
	s := &st{}
	v := reflect.ValueOf(s)
	v.MethodByName("Echo").Call(nil)
	return
}