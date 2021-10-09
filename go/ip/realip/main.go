package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetRealIP() string {
	responseClient, errClient := http.Get("http://ip.dhcp.cn/?ip") // 获取外网 IP
	if errClient != nil {
		fmt.Printf("获取外网 IP 失败，请检查网络\n")
		return ""
	}
	defer responseClient.Body.Close()

	body, _ := ioutil.ReadAll(responseClient.Body)
	clientIP := fmt.Sprintf("%s", string(body))
	return clientIP
}

func LoadRealIP(name string) string {
	fileHandle, fileErr := os.Open(name)
	if nil == fileErr {
		// 获取配置的ip，可以依据需求进行修改
		realIp, realErr := ioutil.ReadAll(fileHandle)
		if nil == realErr {
			return string(realIp)
		}

		webIp := GetRealIP()
		if "" != webIp {
			_ = ioutil.WriteFile(name, []byte(webIp), os.ModePerm)
		}
		return webIp
	}
	webIp := GetRealIP()
	if "" != webIp {
		_ = ioutil.WriteFile(name, []byte(webIp), os.ModePerm)
	}
	return webIp
}

func main() {
	fmt.Println(LoadRealIP("agent"))
	return
}
