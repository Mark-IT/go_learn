package main

import "fmt"

// 实现一个冒泡排序

func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j] < a[j-1] { // 从小到大排序
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
}

func main() {
	b := [...]int{9, 1, 3, 6, 5, 10, 8}
	fmt.Println(b)
	bubbleSort(b[:])
	fmt.Println(b)

}
