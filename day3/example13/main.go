package main

import (
	"fmt"
	"math/rand"
)

// 猜数字，写一个程序，随机生成一个0到100的整数n，然后用户在终端，输入数字，
// 如果和n相等，则提示用户猜对了。如果不相等，则提示用户，大于或小于n。

func main() {
	var n int
	n = rand.Intn(100)
	for {
		var input int
		fmt.Scanf("%d\n", &input) // 因为Scanf默认将换行符视为空格，所以需要输入换行符来匹配格式中的换行
		flag := false
		switch {
		case input == n:
			fmt.Println("you are right")
			flag = true
		case input > n:
			fmt.Println(" bigger")
		default:
			fmt.Println("less")
		}

		if flag {
			return
		}

	}
}
