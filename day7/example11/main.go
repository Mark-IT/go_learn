package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
文件写入
os.OpenFile(“output.dat”,  os.O_WRONLY|os.O_CREATE, 0666)
第二个参数：文件打开模式：
1. os.O_WRONLY：只写
2. os.O_CREATE：创建文件
3. os.O_RDONLY：只读
4.  os.O_RDWR：读写
5.  os.O_TRUNC ：清空

第三个参数：权限控制：
r ——> 004
w——> 002
x——> 001

 */

func main() {
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Println("An error occurred with file creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world \n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}

	outputWriter.Flush() // 必须执行此操作，将数据刷新到硬盘中

}
