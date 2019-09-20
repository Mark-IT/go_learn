package main

import "fmt"

// label 语句形式为  字符串:
func main() {
LABEL1:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if j == 2 {
				continue LABEL1 // 跳到LABEL1处，此处作用相当于 break
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}
