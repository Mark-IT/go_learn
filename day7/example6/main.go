package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("len of args:%d\n", len(os.Args))
	for i, v := range os.Args {	// os.Args是一个string的切片，用来存储所有的命令行参数
		fmt.Printf("args[%d]=%s\n", i, v)
	}
}
