package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	err = c.Send("auth", "123456")
	if err != nil {
		fmt.Println("auth redis failed,", err)
		return
	}

	defer c.Close()

	_, err = c.Do("rpush", "book_list", 100, "567")
	if err != nil {
		fmt.Println(err)
		return
	}

	r1, err := redis.String(c.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	r2, err := redis.String(c.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Printf("%T %v\n", r1, r1)
	fmt.Printf("%T %v\n", r2, r2)
}
