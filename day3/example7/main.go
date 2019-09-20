package main

import "fmt"

// 写一个函数，传入一个int类型的指针，并在函数中修改所指向的值。
// 在main函数中调用这个函数，并把修改前后的值打印到终端，观察结果

func modify(p *int) {
	fmt.Println(p)
	*p = 100
	return
}

func main() {
	a := 1
	var p *int
	p = &a
	fmt.Println(*p)
	modify(p)
	fmt.Println(*p)
	modify(&a)
	fmt.Println(*p)

}
