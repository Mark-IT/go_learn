package main

import "fmt"

// 判断 101-200 之间有多少个素数，并输出所有素数

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true

}

func main() {
	var n int // 区间起始值
	var m int // 区间终点值
	fmt.Scanf("%d%d", &n, &m)
	for i := n; i <= m; i++ {
		if isPrime(i) == true {
			fmt.Printf("[%d] is prime number \n", i)
			continue
		}
	}
}
