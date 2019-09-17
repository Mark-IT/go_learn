package main

import (
	"fmt"
	"os"
	"runtime"
)

// 写一个程序获取当前运行的操作系统名称和PATH环境环境变量的值，并打印在终端

func main() {
	fmt.Printf("The operating system is %s\n", runtime.GOOS)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
