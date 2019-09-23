package main

import "fmt"

/*  struct

1. 用来自定义复杂数据结构
2. struct里面可以包含多个字段（属性）
3. struct类型可以定义方法，注意和函数的区分
4. struct类型是值类型
5. struct类型可以嵌套
6. Go语言没有class类型，只有struct类型

struct中的所有字段在内存是连续的
访问形式如下
stu.Name、stu.Age和stu.Score或者 (*stu).Name、(*stu).Age等

 */

type Student struct {
	Name  string
	Age   int
	score float32
}

func main() {

	var stu Student

	stu.Age = 18
	stu.Name = "hua"
	stu.score = 80

	var stu2 *Student = new(Student)	// 指向结构体的指针

	stu2.Name = "student2"
	stu2.Age = 19
	stu2.score = 90

	var stu1 *Student = &Student{	// 指向结构体的指针
		Age:  20,
		Name: "hua",
	}

	fmt.Println(stu)
	fmt.Println(stu1)
	fmt.Println(stu2)
	fmt.Println(stu.Name)
	fmt.Printf("Name:%p\n", &stu.Name)
	fmt.Printf("Age: %p\n", &stu.Age)
	fmt.Printf("score:%p\n", &stu.score)
}
