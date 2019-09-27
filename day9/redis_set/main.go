package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	options := redis.DialPassword("1234561")
	c, err := redis.Dial("tcp", "localhost:6379", options) // 通过密码登录
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer c.Close()

	_, err = c.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(r)
}
