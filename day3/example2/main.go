package main

import (
	"fmt"
	"strings"
)

/*
strings.Index(s string, str string) int：判断str在s中首次出现的位置，如果没有出现，则返回-1
strings.LastIndex(s string, str string) int：判断str在s中最后出现的位置，如果没有出现，则返回-1
 */

func StrIndex(str string, substr string) (int, int) { // 写一个函数返回一个字符串在另一个字符串的首次出现和最后出现位置
	return strings.Index(str, substr), strings.LastIndex(str, substr)
}

func main() {
	fmt.Println(StrIndex("hello world", "l"))
}
