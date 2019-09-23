package main

import "fmt"

type integer int

type Student struct {
	Number int
}

type Stu Student //alias	给Student取别名为Stu

func main() {

	var i integer = 1000
	var j int = 100

	j = int(i)	// 不能直接 j = i  ，必须强转
	fmt.Println(j)

	var a Student
	a = Student{30}

	var b Stu
	a = Student(b)
	fmt.Println(a)
	fmt.Println(b)

}
