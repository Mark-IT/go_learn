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

函数参数传递方式
1). 值传递
2). 引用传递

注意1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝，引用传递是地址的拷贝。
一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。

注意2：map、slice、chan、指针、interface默认以引用的方式传递

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
