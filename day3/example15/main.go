package main

import "fmt"

func main() {
	str := "hello world,中国" // 英文1个字节编码，中文3个字节编码
	for index, val := range str {	// range 用来遍历数组、slice、map、chan
		fmt.Printf("index[%d] val[%c] len[%d] type[%T] \n", index, val, len(string(val)), val)
	}
}
