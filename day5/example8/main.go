package main

import "fmt"

/*
如果一个struct嵌套了另一个匿名结构体，那么这个结构可以直接访问
匿名结构体的方法，从而实现了继承。

如果一个struct嵌套了另一个有名结构体，那么这个模式就叫组合。

如果一个struct嵌套了多个匿名结构体，那么这个结构可以直接访问
多个匿名结构体的方法，从而实现了多重继承。

 */

type Car struct {
	weight int
	name   string
}

func (p *Car) Run() {
	fmt.Println("running")
}

type Bike struct {
	Car // 继承
	lunzi int
}

type Train struct {
	c Car // 组合
}

func main() {
	var a Bike
	a.weight = 100  // 即 a.Car.weight
	a.name = "bike" // 即 a.Car.name
	a.lunzi = 2

	fmt.Println(a)
	a.Run()

	var b Train
	b.c.weight = 100
	b.c.name = "train"
	b.c.Run()
}
