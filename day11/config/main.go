package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("ini", "D:\\go_test\\src\\go_learn\\day11\\config\\logagent.conf")	// 按 ini格式读取配置
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("read server:port failed, err:", err)
		return
	}

	fmt.Println("Port:", port)
	logLevel := conf.String("logs::log_level")
	if len(logLevel) == 0 {
		logLevel = "debug"
	}

	fmt.Println("log_level:", logLevel)

	logPath := conf.String("logs::log_path")
	fmt.Println("log_path:", logPath)
}
