package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//  从终端读取一行字符串，统计英文、数字、空格以及其他字符的数量。

type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	file, err := os.Open("C:/test.log")
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	defer file.Close()

	var count CharCount

	reader := bufio.NewReader(file)		// 带缓冲区的文件读写
	for {
		str, err := reader.ReadString('\n')	// 按行读取
		if err == io.EOF {	// 文件读取完毕
			break
		}
		if err != nil {
			fmt.Printf("read file failed, err:%v", err)
			break
		}

		runeArr := []rune(str)
		for _, v := range runeArr {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}

	}

	fmt.Printf("char count:%d\n", count.ChCount)
	fmt.Printf("num count:%d\n", count.NumCount)
	fmt.Printf("space count:%d\n", count.SpaceCount)
	fmt.Printf("other count:%d\n", count.OtherCount)
}
