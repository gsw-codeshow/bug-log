package main

import (
	"errors"
	"fmt"
	"net"
	"os"
)

func getIp() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}

func GetLocalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if nil != err {
		return nil, err
	}
	for _, iface := range ifaces {
		if 0 == iface.Flags&net.FlagUp {
			continue
		}
		if 0 != iface.Flags&net.FlagLoopback {
			continue
		}
		addrs, err := iface.Addrs()
		if nil != err {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if nil == ip {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if nil == ip || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	return ip
}

func main() {
	ip, err := externalIP()
	if nil != err {
		fmt.Println(err)
	}
	fmt.Println(ip.String())
	return
}
