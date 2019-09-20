package main

import "fmt"

// 输入一个字符串，判断其是否为回文。回文字符串是指从左到右读和从右到左读完全相同的字符串。

/*
func process(s string) bool {	// 不能处理中文,
    length := len(s)
	for i := 0; i < len(s); i++ {	// 遍历时，处理的是字符，中文占3个字符
		if i == length/2 {
			break
		}
		last := length - i - 1
		if s[last] != s[i] {
			return false
		}
	}
	return true
}
*/

func process(s string) bool { // 能处理中文
	t := []rune(s) // rune 处理字符，中英文均可
	length := len(t)
	for i, _ := range t {
		if i == length/2 {
			break
		}
		last := length - i - 1
		if t[last] != t[i] {
			return false

		}
	}
	return true

}

func main() {
	var str string
	fmt.Scanf("%s", &str)
	if process(str) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
