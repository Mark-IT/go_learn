package main

import (
	"fmt"
	"strconv"
)

// 写一个程序，从终端读取输入，并转成整数，如果转成整数出错，则输出 “can not convert to int”，并返回。否则输出该整数。

func main() {
	var str string
	fmt.Scanf("%s", &str)
	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("%s can not convert to int", str)
		return
	} else {
		fmt.Println(number)
	}
}
