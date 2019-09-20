package main

import (
	"bufio"
	"fmt"
	"os"
)

// 输入一行字符，分别统计出 英文字母、空格、数字和其它字符的个数

func count(str string) (worldCount, spaceCount, numberCount, otherCount int) {
	t := []rune(str) // rune 处理字符
	for _, v := range t {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			worldCount ++
		case v == ' ':
			spaceCount++
		case v >= '0' && v <= '9':
			numberCount++
		default:
			otherCount++
		}

	}
	return

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err:", err)
	}
	wc, sc, nc, oc := count(string(result))
	fmt.Printf(" word count:%d\n space count:%d\n number count:%d\n other count:%d\n", wc, sc, nc, oc)
}
