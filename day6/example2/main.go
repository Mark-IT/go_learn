package main

import "fmt"

type Carer interface {
	GetName() string
	Run()
	DiDi()
}

type Test interface {
	Hello()
}

type BMW struct {
	Name string
}

func (p *BMW) GetName() string {
	return p.Name
}

func (p *BMW) Run() {
	fmt.Printf("%s is running\n", p.Name)
}

func (p *BMW) DiDi() {
	fmt.Printf("%s is didi\n", p.Name)
}

func (p *BMW) Hello() {
	fmt.Printf("hello, i'm %s \n", p.Name)
}


type BYD struct {
	Name string
}

func (p *BYD) GetName() string {
	return p.Name
}

func (p *BYD) Run() {
	fmt.Printf("%s is running\n", p.Name)
}

func (p *BYD) DiDi() {
	fmt.Printf("%s is didi\n", p.Name)
}


func main() {
	var car Carer	// 接口类型默认是指针
	var test Test	// 接口类型默认是指针

	fmt.Println(car)

	bmw := &BMW{
		Name: "BMW",
	}

	car = bmw
	car.Run()
	bmw.Run()

	test = bmw
	test.Hello()

	byd := &BYD{
		Name: "BYD",
	}

	car = byd
	car.Run()
}
