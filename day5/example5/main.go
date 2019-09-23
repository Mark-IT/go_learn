package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"student_name"`	// 定义tag,json序列化时打印的key，没定义时打印大写的字段名
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

type Student1 struct {
	Name  string `json:"student_name"`
	age   int    `json:"age"`	// 小写的不会被打印出来
	score int    `json:"score"`	// 小写的不会被打印出来
}
func main() {
	var stu Student = Student{
		Name:  "stu",
		Age:   18,
		Score: 80,
	}

	var stu1 Student1 = Student1{
		Name:  "stu1",
		age:   18,
		score: 80,
	}

	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json encode stu failed, err:", err)
		return
	}

	//fmt.Println(data)	 // 字符数组，ASCII码
	fmt.Println(string(data))

	data1, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("json encode stu failed, err:", err)
		return
	}
	fmt.Println(string(data1))
}
