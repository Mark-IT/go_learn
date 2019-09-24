package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

/*

用反射操作结构体

a. reflect.Value.NumField()获取结构体中字段的个数
a. reflect.Value.NumMethod()获取结构体中方法的个数
b. reflect.Value.Method(n).Call来调用结构体中的方法

 */

type Student struct {
	Name  string `json:"student_name"`
	Age   int    `json:"age"`
	Score float32
	Sex   string
}

func (s Student) Print() {
	fmt.Println("---start----")
	fmt.Println(s)
	fmt.Println("---end----")
}

func (s Student) Set(name string, age int, score float32, sex string) {

	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	tye := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := val.Elem().NumField()             //  获取结构体中字段的个数
	val.Elem().Field(0).SetString("stu1000") // 修改第一个字段的值
	for i := 0; i < num; i++ {
		fmt.Printf("%d \t %v \t %v\n", i, val.Elem().Field(i), val.Elem().Field(i).Kind())
	}

	fmt.Printf("struct has %d fields\n", num)

	tag := tye.Elem().Field(0).Tag.Get("json") // 获取第一个元素的 json tag
	fmt.Printf("tag=%s\n", tag)

	numOfMethod := val.Elem().NumMethod() // 获取结构体中方法的个数
	fmt.Printf("struct has %d methods\n", numOfMethod)
	var params []reflect.Value
	val.Elem().Method(0).Call(params) // 调用第一个方法（按定义的顺序）
}

func main() {
	var a Student = Student{
		Name:  "stu01",
		Age:   18,
		Score: 92.8,
	}

	result, _ := json.Marshal(a)
	fmt.Println("json result:", string(result))

	TestStruct(&a)
	fmt.Println(a)
}
