package main

import (
	"flag" // 命令行解析包
	"fmt"
)

func main() {
	var confPath string
	var logLevel int
	flag.StringVar(&confPath, "c", "", "please input conf path") // 将命令行中c 后面的参数值赋予confPath变量，默认值为value对应的值 ,string类型
	flag.IntVar(&logLevel, "d", 10, "please input log level")    // 将命令行中d 后面的参数值赋予logLevel变量，默认值为value对应的值，int类型

	flag.Parse()

	fmt.Println("path:", confPath)
	fmt.Println("log level:", logLevel)
}
