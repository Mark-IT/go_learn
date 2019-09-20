package main

import (
	"fmt"
	"strings"
)

func Print(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(strings.Repeat("A", i))
	}
}
func main() {
	Print(5)
}
