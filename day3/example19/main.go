package main

import "fmt"

type myFunc func(int, int) int // 定义一个myFunc 函数类型

func add(a, b int) int {
	return a + b
}

func operator(op myFunc, a int, b int) int {
	return op(a, b)
}

func abc(a1, b1, c1 int) (a, b, c int) { // 命名返回值的名字
	c = a1
	b = b1
	a = c1
	return
	//return a, b, c
}

func main() {
	sum := add
	fmt.Println(sum)
	sum1 := operator(sum, 100, 200)
	fmt.Println(sum1)
	a := 1
	b := 2
	c := 3

	a2, b2, c2 := abc(a, b, c)
	fmt.Println(a2, b2, c2)

	a3, _, c3 := abc(a2, b2, c2) //  _标识符，用来忽略返回值
	fmt.Println(a3, c3)
}
