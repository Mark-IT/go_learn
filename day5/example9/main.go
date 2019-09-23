package main

import "fmt"

type Car struct {
	weight int
	name   string
}

func (p *Car) Run() {
	fmt.Println("running")
}

type Bike struct {
	Car
	lunzi int
}

type Train struct {
	Car
}

// 如果一个变量x实现了String()这个方法，那么fmt.Printf(%s,&x)默认会调用这个变量的String()进行输出。

func (p *Train) String() string {
	str := fmt.Sprintf("name=[%s] weight=[%d]", p.name, p.weight)
	return str
}

func main() {
	var a Bike
	a.weight = 100
	a.name = "bike"
	a.lunzi = 2

	fmt.Println(a)
	a.Run()

	var b Train
	b.weight = 100
	b.name = "train"
	b.Run()

	fmt.Printf("%s", &b)
}
