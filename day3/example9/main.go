package main

import "fmt"

func main() {
	a := 0

	switch a {
	case 0:
		fmt.Println("a is equal 0")	// 不需要加break,满足条件后不会往后匹配
		fmt.Println("yes")
	case 10:
		fmt.Println("a is equal 10")
	default:
		fmt.Println("a is equal default")

	}
}
