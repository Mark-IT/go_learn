package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("C:/test.log")
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	defer file.Close()

}
