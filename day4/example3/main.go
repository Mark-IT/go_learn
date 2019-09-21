package main

import (
	"fmt"
)

func recursion(n int) {
	fmt.Println("hello", n)
	if n > 10 {
		return
	}
	recursion(n + 1)

}

func factor(n int) int { // 阶层
	if n == 1 {
		return 1
	}
	return factor(n-1) * n
}

func fab(n int) int { // 斐波那契数
	if n <= 1 {
		return 1
	}
	return fab(n-1) + fab(n-2)
}

func main() {
	recursion(0)
	fmt.Println(factor(5))
	for i := 0; i < 10; i ++ {
		fmt.Println(fab(i))

	}
}
