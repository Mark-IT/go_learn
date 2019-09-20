package main

import "fmt"

// 编写程序，在终端输出九九乘法表。

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
