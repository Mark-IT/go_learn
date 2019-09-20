package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {

	c := add //  函数是一等公民，函数也是一种类型，一个函数可以赋值给变量

	fmt.Println(c)

	sum := c(10, 20)
	fmt.Println(sum)

}
