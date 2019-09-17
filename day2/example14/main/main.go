package main

import (
	"fmt"
	"math/rand"
	"time"
)

//使用math/rand生成10个随机整数，10个小于100的随机整数以及10个随机浮点数
func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {

	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)
	}

	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}

	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}
}
