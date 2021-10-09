package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

var randIndex = rand.New(rand.NewSource(time.Now().Unix()))

func RandInt(max int) int {
	return randIndex.Intn(max)
}

func GetInternalIP() (string, error) {
	// 思路来自于Python版本的内网IP获取，其他版本不准确
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", errors.New("internal IP fetch failed, detail:" + err.Error())
	}
	defer conn.Close()

	// udp 面向无连接，所以这些东西只在你本地捣鼓
	res := conn.LocalAddr().String()
	res = strings.Split(res, ":")[0]
	return res, nil
}

func GetExternalIP() (string, error) {
	// 有很多类似网站提供这种服务，这是我知道且正在用的
	// 备用：https://myexternalip.com/raw （cip.cc 应该是够快了，我连手机热点的时候不太稳，其他自己查）
	response, err := http.Get("http://ip.cip.cc")
	if err != nil {
		return "", errors.New("external IP fetch failed, detail:" + err.Error())
	}

	defer response.Body.Close()
	res := ""

	// 类似的API应当返回一个纯净的IP地址
	for {
		tmp := make([]byte, 32)
		n, err := response.Body.Read(tmp)
		if err != nil {
			if err != io.EOF {
				return "", errors.New("external IP fetch failed, detail:" + err.Error())
			}
			res += string(tmp[:n])
			break
		}
		res += string(tmp[:n])
	}

	return strings.TrimSpace(res), nil
}

func main() {
	ip, err := GetInternalIP()
	fmt.Println(" --- --- GetInternalIP --- --- ")
	fmt.Println(ip)
	fmt.Println(" --- --- --- ")
	fmt.Println(err)
	type UnStruct struct {
		Name string
	}
	var i interface{}
	i = int64(1)
	test, ok := i.(UnStruct)
	fmt.Println(ok)
	fmt.Println(test)
	return
}
