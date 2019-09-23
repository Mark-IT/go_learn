package main

import (
	"fmt"
	"math/rand"
)

// 单链表

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student
}

func traverse(p *Student) { // 遍历链表
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func insertTail(p *Student) { // 尾插
	var tail = p

	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}

		tail.next = &stu
		tail = &stu

	}
}

func insertHead(p **Student) { // 头插
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		stu.next = *p
		*p = &stu
	}
}

func addNode(p *Student, newNode *Student) {
	for p != nil {
		if p.Name == "stu9" {
			newNode.next = p.next
			p.next = newNode
			break
		}
		p = p.next //移动指针
	}
}

func delNode(p *Student) {
	var prev *Student = p
	for p != nil {
		if p.Name == "stu6" { //删除name为 stu6的节点
			prev.next = p.next
			break
		} else {
			prev = p
			p = p.next
		}

	}
}

func main() {

	// head := Student{
	//	Name:  "hua",
	//	Age:   18,
	//	Score: 100,
	//}
	var head *Student = new(Student)

	head.Name = "hua"
	head.Age = 18
	head.Score = 100

	//traverse(head)

	insertTail(head)
	//traverse(head)

	//insertHead(&head)
	//traverse(head)

	delNode(head)
	traverse(head)

	fmt.Println()

	var newNode *Student = new(Student)

	newNode.Name = "stu1000"
	newNode.Age = 18
	newNode.Score = 100
	addNode(head, newNode)

	traverse(head)

}
