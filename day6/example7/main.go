package main

import (
	"fmt"
	"reflect"
)

/*
反射：可以在运行时动态获取变量的相关信息
     Import (“reflect”)

a. reflect.TypeOf，获取变量的类型，返回reflect.Type类型
b. reflect.ValueOf，获取变量的值，返回reflect.Value类型
c. reflect.Value.Kind，获取变量的类别，返回一个常量
d. reflect.Value.Interface()，转换成interface{}类型

获取变量的值：
reflect.ValueOf(x).Float()
reflect.ValueOf(x).Int()
reflect.ValueOf(x).String()
reflect.ValueOf(x).Bool()

通过反射的来改变变量的值

reflect.Value.SetXX相关方法，比如:
reflect.Value.SetFloat()，设置浮点数
reflect.Value.SetInt()，设置整数
reflect.Value.SetString()，设置字符串


 */
type Student struct {
	Name  string
	Age   int
	Score float32
}

func test(b interface{}) {
	t := reflect.TypeOf(b)
	fmt.Println(t)	// main.Student

	v := reflect.ValueOf(b)
	k := v.Kind()
	fmt.Println(k)	// struct

	iv := v.Interface()
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("%v %T\n", stu, stu)
	}
}

func testInt(b interface{}) {
	val := reflect.ValueOf(b)	// val是b的一个拷贝，修改val，b不会修改!


	val.Elem().SetInt(100)	// 解决方法，传地址 val.Elem()用来获取指针指向的变量，相当于：
								// var a *int;
								// *a = 100

	c := val.Elem().Int()
	fmt.Printf("get value  interface{} %d\n", c)
	fmt.Printf("string val:%d\n", val.Elem().Int())

}

func main() {
	var a Student = Student{
		Name:  "stu01",
		Age:   18,
		Score: 92,
	}
	test(a)

	var b int = 1
	b = 200
	testInt(&b)
	fmt.Println(b)

}
