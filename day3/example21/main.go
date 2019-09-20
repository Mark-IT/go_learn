package main

import "fmt"

/*
可变参数
func add(arg…int) int {	0个或多个参数

}

func add(a int, arg…int) int {	   1个或多个参数

}

func add(a int, b int, arg…int) int {	  2个或多个参数

}

注意：其中arg是一个slice，我们可以通过arg[index]依次访问所有参数
通过len(arg)来判断传递参数的个数

 */

// 写一个函数add，支持1个或多个int相加，并返回相加结果

func add(a int, arg ...int) int {
	var sum int = a
	for i := 0; i < len(arg); i++ {
		sum += arg[i]

	}
	return sum
}

func main() {
	sum := add(10)
	fmt.Println(sum)

	sum1 := add(10, 1, 2, 3)
	fmt.Println(sum1)
}
