package main

import "fmt"

func main() {
	a := 0
	switch a {
	case 0:
		// a 为0 ，则什么也不做
	case 1, 2: // a 为1或2时，都执行下面的代码
		fmt.Println("1 or 2")

	default:
		fmt.Println(" default")

	}
}
