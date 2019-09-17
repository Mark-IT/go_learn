package main

import (
	"fmt"
	"go_learn/day2/example3/add"
)

func main() {
	fmt.Println("Name:",add.Name)	// ""
	fmt.Println("age:",add.Age)		// 0
	add.Test()
	fmt.Println("Name:",add.Name)	// "hello world"
	fmt.Println("age:",add.Age)		// 1
}
