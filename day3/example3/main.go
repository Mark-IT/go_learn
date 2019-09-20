package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*

strings.Replace(str string, old string, new string, n int)：字符串替换
strings.Count(str string, substr string)int：字符串计数
strings.Repeat(str string, count int)string：重复count次str
strings.ToLower(str string)string：转为小写
strings.ToUpper(str string)string：转为大写
strings.TrimSpace(str string)：去掉字符串首尾空白(含\n)字符
strings.Trim(str string, cut string)：去掉字符串首尾cut字符
strings.TrimLeft(str string, cut string)：去掉字符串首cut字符
strings.TrimRight(str string, cut string)：去掉字符串首cut字符
strings.Fields(str string)：返回str空格分隔的所有子串的slice
strings.Split(str string, split string)：返回str split分隔的所有子串的slice
strings.Join(s1 []string, sep string)：用sep把s1中的所有元素链接起来
strconv.Itoa(i int)：把一个整数i转成字符串
strconv.Atoi(str string)(int, error)：把一个字符串转成整数

 */

func main() {
	s := "   Hello World   \n"

	result := strings.Replace(s, "l", "x", -1) // -1 代表全替换，0不替换
	fmt.Println("replace:", result)

	count := strings.Count(s, "l")
	fmt.Println("count:", count)

	result = strings.Repeat(s, 3)
	fmt.Println("repeat:", result)

	upperS := strings.ToUpper(s)
	fmt.Println("upper:", upperS)

	lowerS := strings.ToLower(upperS)
	fmt.Println("lower:", lowerS)

	result = strings.TrimSpace(s)
	fmt.Println("trimSpace:", result)

	result = strings.Trim(s, "\n\r")
	fmt.Println("trim:", result)

	result = strings.TrimLeft(s, "\n\r")
	fmt.Println("trimLeft:", result)

	result = strings.TrimRight(s, "\n\r")
	fmt.Println("trimRight:", result)

	splitResult := strings.Fields(s)
	fmt.Println("Fields:")
	for i := 0; i < len(splitResult); i++ {
		fmt.Println(splitResult[i])
	}

	splitResult = strings.Split(s, "l")
	fmt.Println("split:")
	for i := 0; i < len(splitResult); i++ {
		fmt.Println(splitResult[i])
	}

	str2 := strings.Join(splitResult, "l")
	fmt.Println("Join:", str2)

	str3 := strconv.Itoa(1000)
	fmt.Println("itoa:", str3)

	number, err := strconv.Atoi(str3)
	if err != nil {
		fmt.Println("can not convert to int", err)
		return
	} else {
		fmt.Println("number:", number)
	}

}
