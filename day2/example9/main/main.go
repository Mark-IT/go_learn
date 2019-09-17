package main

import "fmt"

//写一个程序，交换两个整数的值。比如： a=3; b=4; 交换之后：a=4;b=3

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func swap1(a int, b int) (int, int) {
	return b, a
}

func main() {
	first := 100
	second := 200

	// 方法一
	//swap(&first,&second)

	// 方法二
	// first, second = swap1(first, second)

	// 方法三
	first, second = second, first

	fmt.Println("first=", first)
	fmt.Println("second=", second)

}
