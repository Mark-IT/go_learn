package main

import (
	"fmt"
	"time"
)

type Car struct {
	Name string
	Age  int
}

func (c *Car) Set(name string, age int) {
	c.Name = name
	c.Age = age
}

type Car2 struct {
	Name string
}

type Train struct {
	Car
	Car2
	createTime time.Time
	int
}

func (t *Train) Set(age int) {
	t.int = age
}

func main() {
	var train Train
	train.int = 300
	fmt.Println(train)
	train.Set(100) // 等价于(&train).Set(100)
	fmt.Println(train)
	train.Car.Set("huas", 1001)
	train.Car.Name = "test"
	fmt.Println(train)

}
