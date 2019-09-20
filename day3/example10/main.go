package main

import "fmt"

func main() {
	a := 0

	switch a {
	case 0:
		fmt.Println("0")
		fallthrough	//  条件匹配成功后，如果含义fallthrough，则会继续往下匹配一次
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	default:
		fmt.Println(" default")

	}
}
