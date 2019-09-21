package main

import "fmt"

func test1() {
	var a [10]int	// 数组 一旦定义，长度不能变，长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型
	a[0] = 100
	a[9] = 200

	fmt.Println("test1 a:", a)

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for index, val := range a {
		fmt.Printf("a[%d]=%d\n", index, val)
	}
}

func test3(arr *[5]int) {
	(*arr)[0] = 1000
}

func test2() {
	var a [10]int
	b := a

	b[0] = 100 //  数组是值类型，因此改变副本的值，不会改变本身的值

	fmt.Println("test2 a:", a)
}

func main() {

	test1()
	test2()
	var a [5]int
	test3(&a)
	fmt.Println("test3 a:", a)
}
