package main

import "fmt"

/*
字符串表示两种方式： 1）双引号    2）``   （反引号）
反引号里的内容不会转义，会原封不动的输出
 */

func main() {
	var str = "hello world \n\n"
	var str2 = `hello world \n\n
this is a test string
This is a test string too
`
	fmt.Println("str=",str)
	fmt.Println("str2=",str2)



}