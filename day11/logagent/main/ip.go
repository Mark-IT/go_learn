package main

import (
	"fmt"
	"net"
)

var (
	localIPArray []string
)

func init() {
	addrs, err := net.InterfaceAddrs() // 获取所有网卡的地址
	if err != nil {
		panic(fmt.Sprintf("get local ip failed, %v", err))
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() { // 排除回环地址
			if ipnet.IP.To4() != nil {
				localIPArray = append(localIPArray, ipnet.IP.String())
			}
		}
	}

	fmt.Println(localIPArray)
}
