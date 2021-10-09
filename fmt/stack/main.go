package main

import (
	"fmt"
	"runtime/debug"

	"github.com/pkg/errors"
)

func test1() {
	test2()
}

func test2() {
	test3()
}

func test3() {
	// 可以通过 debug.PrintStack() 直接打印，也可以通过 debug.Stack() 方法获取堆栈然后自己打印
	fmt.Printf("%s", debug.Stack())
	debug.PrintStack()
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}
type StackTrace []errors.Frame

func main() {
	//test1()
	err := errors.New(" --- --- ")
	errors.Wrap(err, "read failed")

	fmt.Println(" --- --- --- 2 --- --- --- ")
	err1 := errors.New(" --22- --- ")

	errors.Wrap(err1, "0000")
	if err, ok := err.(stackTracer); ok {
		for _, f := range err.StackTrace() {
			fmt.Printf("%+s:%d\n", f, f)
		}
	}
	return
}
