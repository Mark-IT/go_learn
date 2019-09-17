package main

import "fmt"

// 写一个程序用来打印值类型和引用类型变量到终端，并观察输出结果

/*
值类型：
	变量直接存储值，内存通常在栈中分配。
	基本数据类型int、float、bool、string以及数组和struct。

引用类型：
	变量存储的是一个地址，这个地址存储最终的值。内存通常在堆上分配。通过GC回收。
	指针、slice、map、chan等都是引用类型。

 */

func modify(a int) {
	// 值传递，修改的是拷贝的值，对原值没影响
	a = 10
	return
}

func modify1(a *int) {
	// 引用地址传递，修改的是所指向的地址的值，即会修改传入的原值
	*a = 10		// 这里写成 a = 10 是不行的， 这里的a是地址的副本，写成 a=10,是修改的地址副本的值
				// *a 指操作这个副本所指向的内存地址的值
}

func main() {
	a := 5
	b := make(chan int, 1)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	modify(a)
	fmt.Println("modify a=", a)
	modify1(&a)
	fmt.Println("modify1 a=", a)

}
