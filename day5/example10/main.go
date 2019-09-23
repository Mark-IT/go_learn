package main

import "fmt"

type People struct {
	name string
	age  int
}

/*
interface类型默认是一个指针

interface类型可以定义一组方法，但是这些不需要实现。并且interface不能包含任何变量。

Golang中的接口，不需要显式的实现。只要一个变量，含有接口类型中的所有方法，那么这个变量就实现这个接口
如果一个变量含有了多个interface类型的方法，那么这个变量就实现了多个接口。
如果一个变量只含有了1个interface的方部分方法，那么这个变量没有实现这个接口

一个接口可以嵌套在另外的接口

 */

type Test interface {
	Print()
	Sleep()
}

type Student struct {
	name  string
	age   int
	score int
}

func (p Student) Print() {
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
	fmt.Println("score:", p.score)
}

func (p Student) Sleep() {
	fmt.Println("student sleep")
}

func (people People) Print() {
	fmt.Println("name:", people.name)
	fmt.Println("age:", people.age)
}

func (p People) Sleep() {
	fmt.Println("people sleep")
}

func main() {

	var t Test
	fmt.Println(t)
	//t.Print()

	var stu Student = Student{
		name:  "stu1",
		age:   20,
		score: 200,
	}

	t = stu
	t.Print()
	t.Sleep()
	fmt.Println()
	var people People = People{
		name: "people",
		age:  100,
	}

	t = people
	t.Print()
	t.Sleep()

	fmt.Println("t:", t)

	fmt.Println()

	// 类型断言，由于接口是一般类型，不知道具体类型，如果要转成具体类型可以采用以下方法进行转换：

	n := "hello world"
	var x interface{} // 空接口没有任何方法，所以所有类型都实现了空接口。

	x = n
	y, ok := x.(int) // 通过判断x是否能转成为int来，判断
	if !ok {
		fmt.Println("n is not int")

	} else {
		fmt.Print(y)

	}

}
