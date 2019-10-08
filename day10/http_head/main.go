package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

var url = []string{
	"http://www.baidu.com",
	"http://google.com",
	"http://taobao.com",
}

func main() {
	for _, v := range url {
		c := http.Client{
			Transport: &http.Transport {
				DialContext:func(ctx context.Context,network, addr string) (net.Conn, error){
					timeout := time.Second*2	// 设置 tcp 连接阶段的超时
					return net.DialTimeout(network, addr, timeout)
				},
			},
		}
		resp, err := c.Head(v)	// 发送head请求
		if err != nil {
			fmt.Printf("head %s failed, err:%v\n", v, err)
			continue
		}

		fmt.Printf("head succ, status:%v\n", resp.Status)
	}
}
