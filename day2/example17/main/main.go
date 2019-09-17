package main

import "fmt"

// 字符串反转

func reverse(str string) string {
	var result string
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		result += fmt.Sprintf("%c", str[strLen-i-1])
	}
	return result

}

func reverse1(str string) string {
	var result []byte
	tmp := []byte(str)
	length := len(str)
	for i := 0; i < length; i++ {
		result = append(result, tmp[length-i-1])
	}
	return string(result)

}
func main() {
	str1 := "hello"
	str2 := "world"
	str3 := str1 + " " + str2
	fmt.Println(str3)
	str4 := reverse(str3)
	fmt.Println(str4)
	str5 := reverse1(str4)
	fmt.Println(str5)

}
