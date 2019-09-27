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

	_, err = c.Do("Hset", "books", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("HGet", "books", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Printf("%T %v\n", r, r)
}
