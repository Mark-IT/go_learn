package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//options := redis.DialPassword("1234561")
	//c, err := redis.Dial("tcp", "localhost:6379", options)	// 通过密码登录，则不用发送auth
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	// 如果redis 有密码 ，在连接成功后必须认证认证的代码，否则后面操作会失败
	err = c.Send("auth", "123456")
	if err != nil {
		fmt.Println("auth redis failed,", err)
		return
	}

	defer c.Close()
}
