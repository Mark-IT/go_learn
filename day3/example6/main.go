package main

import "fmt"

// 写一个程序，获取一个变量的地址，并打印到终端
func main() {
	a := 10
	fmt.Println("a的地址：", &a) // &取变量的地址

	var p *int               //	声明一个Int型的指针
	p = &a                   // 给指针类型赋值，让其指向一个内存地址
	fmt.Println("p的值为：", *p) // 打印指针所指向的内存地址的值

	*p = 100 // 把 *p所指向的内存地址的值改成了100，即a的值变成了100
	fmt.Println("p的值为:", *p)
	fmt.Println("a的值为:", a)

	b := 999
	p = &b
	*p = 5
	fmt.Println("p的值为:", *p)
	fmt.Println("a的值为:", a)
	fmt.Println("b的值为:", b)

}
