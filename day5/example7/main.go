package main

import "fmt"

type integer int

func (p integer) print() {	// 给integer类型自定义方法
	fmt.Println("p is ", p)
}

func (p *integer) set(b integer) {	// 要进行设置，需要传入指针
	*p = b
}

type Student struct {
	Name  string
	Age   int
	Score int
	sex   int
}

func (p *Student) init(name string, age int, score int) {	// 给Student类型自定义方法
	p.Name = name
	p.Age = age
	p.Score = score
	fmt.Println(p)
}

func (p Student) get() Student {
	return p
}

func main() {
	var stu Student
	stu.init("stu", 10, 200)

	stu1 := stu.get()
	fmt.Println(stu1)

	var a integer
	a = 100
	a.print()

	a.set(1000)
	a.print()
}
