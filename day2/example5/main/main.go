package main

import (
	"fmt"
	a "go_learn/day2/example3/add" // 给包定义一个别名a,则可以通过a点来调用原包中的变量或函数
)

//当我们导入一个包，但是不引用里面的任何变量或函数时，只是做一个初始化时，必须将包的别名定义为_, 否则编译错误

func main() {
	fmt.Println("Name:", a.Name) // "hello world"
	fmt.Println("age:", a.Age)   // 1
}
