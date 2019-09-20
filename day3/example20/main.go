package main

import "fmt"

/*
函数参数传递方式
1). 值传递
2). 引用传递

注意1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝，引用传递是地址的拷贝。
一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。

注意2：map、slice、chan、指针、interface默认以引用的方式传递

 */

var (
	result = func(a1 int, b1 int) int { // 全局定义了一个匿名函数
		return a1 + b1
	}
)

func modify(a int) {
	a = 100
}

func test(a, b int) int {
	result := func(a1, b1 int) int { // 定义了一个匿名函数,定义时就调用了
		return a1 + b1
	}(a, b)
	return result
	//result := func(a1, b1 int) int {	// 定义了一个匿名函数
	//	return a1 + b1
	//}
	//return result(a,b)

}

func main() {
	a := 8
	fmt.Println(a)
	modify(a)
	fmt.Println(a)

	fmt.Println(test(100, 200))
	fmt.Println(result(100, 200))
}
