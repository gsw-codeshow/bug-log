package main

import (
	"fmt"
	"net"
	"os"
	"unsafe"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error !", err.Error())
		os.Exit(1)
	}
}

func typeof(v interface{}) {
	fmt.Printf("type is:%T\n", v)
}
func sizeof(v interface{}) {
	fmt.Println("sizeof is: ", unsafe.Sizeof(v))
}

func main() {
	addr := "www.baidu.com:80"
	tcpaddr, err := net.ResolveTCPAddr("", addr)
	checkError(err)

	fmt.Println("tcpaddr is:", tcpaddr)
	fmt.Println("IP is:", tcpaddr.IP.String(), "Port is", tcpaddr.Port)
	typeof(addr)
	typeof(tcpaddr)
	sizeof(addr)
	sizeof(tcpaddr)
	fmt.Println("addr len is:", len(addr))
	fmt.Println("tcpaddr len is:", len(tcpaddr.String()))
	return
}
