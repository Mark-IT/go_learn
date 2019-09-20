package main

import "fmt"

// 写一个函数concat，支持1个或多个string相拼接，并返回结果

func concat(s string, arg ...string) string {
	result := s
	for i := 0; i < len(arg); i++ {
		result += arg[i]
	}
	return result
}

func main() {
	res := concat("hello", " ", "world")
	fmt.Println(res)
}
